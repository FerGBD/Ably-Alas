const ablyApiKey = document.body.getAttribute('data-ably-key');
const ably = new Ably.Realtime(ablyApiKey);
const channel = ably.channels.get('flamengo-vs-vasco');
let lanceCount = 0;

function updateClock() {
    const now = new Date();
    document.getElementById('current-time').textContent = 
        now.toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' });
}

const iconMap = {
    'gol': '⚽',
    'falta': '🚫',
    'cartao-amarelo': '🟨',
    'cartao-vermelho': '🟥',
    'substituicao': '🔄',
    'escanteio': '🚩',
    'penalti': '⭕',
    'impedimento': '🚩',
    'lance': '📢'
};

function getTipoClass(tipo) {
    return tipo.toLowerCase().replace(/\s+/g, '-');
}

function getIcon(tipo) {
    const tipoLower = tipo.toLowerCase();
    return iconMap[tipoLower] || '📢';
}

function formatTimestamp() {
    const now = new Date();
    return now.toLocaleTimeString('pt-BR', { 
        hour: '2-digit', 
        minute: '2-digit',
        second: '2-digit'
    });
}

function processarLance(msg) {
    const status = document.getElementById('status');
    const container = document.getElementById('lances-container');
    const countBadge = document.querySelector('.lance-count');
    
    status.style.display = 'none';

    console.log("Dados recebidos:", msg.data);

    let lance;
    if (typeof msg.data === 'string') {
        try {
            lance = JSON.parse(msg.data);
        } catch (e) {
            lance = { tipo: 'lance', descricao: msg.data, minuto: 0 };
        }
    } else {
        lance = msg.data || { tipo: 'lance', descricao: 'Sem descrição', minuto: 0 };
    }

    const tipoClass = getTipoClass(lance.tipo || 'lance');
    const icon = getIcon(lance.tipo || 'lance');

    const div = document.createElement('div');
    div.className = `lance ${tipoClass}`;
    
    div.innerHTML = `
        <div class="lance-header">
            <span class="icon">${icon}</span>
            <span class="tempo">${lance.minuto || 0}'</span>
            <span class="tipo-badge ${tipoClass}">${lance.tipo || 'Lance'}</span>
            <span class="timestamp">${formatTimestamp()}</span>
        </div>
        <div class="descricao">${lance.descricao || 'Sem descrição'}</div>
        ${lance.autor ? `<div class="autor"><span>✍️</span>${lance.autor}</div>` : ''}
    `;
    
    container.prepend(div);
    
    lanceCount++;
    countBadge.textContent = `${lanceCount} ${lanceCount === 1 ? 'lance' : 'lances'}`;
}

function init() {
    updateClock();
    setInterval(updateClock, 1000);
    
    channel.subscribe('lance', processarLance);
    
    ably.connection.on('connected', () => {
        console.log(' Conectado ao Ably!');
    });
    
    ably.connection.on('failed', () => {
        document.getElementById('status').innerHTML = 
            '<div class="status-icon">❌</div><div>Falha na conexão. Tente recarregar a página.</div>';
    });
}

if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
} else {
    init();
}
