// Analytics simple para English At Lima
document.addEventListener('DOMContentLoaded', () => {
    const waButton = document.getElementById('whatsapp-btn');

    if (waButton) {
        waButton.addEventListener('click', () => {
            // Log en consola para pruebas
            console.log('✅ Evento: Clic en WhatsApp detectado');

            // Esto envía el evento automáticamente a Umami
            if (window.umami) {
                umami.track('Contact-WhatsApp');
            }
        });
    }

    // Medir tiempo de permanencia simple 
    setTimeout(() => {
        console.log('⏱️ Usuario interesado: +30 segundos en la página');
    }, 30000);
});