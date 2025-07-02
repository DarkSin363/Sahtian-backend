const API_BASE_URL = 'http://localhost:11080/api/v1'; // Локальный сервер

document.addEventListener('DOMContentLoaded', async () => {
    await loadClients();
    
    document.getElementById('client-form')?.addEventListener('submit', async (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const client = {
            first_name: formData.get('first_name'),
            last_name: formData.get('last_name'),
            email: formData.get('email')
        };
        
        try {
            const response = await fetch(`${API_BASE_URL}/clients`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(client)
            });
            
            if (response.ok) await loadClients();
        } catch (error) {
            console.error('Error:', error);
        }
    });
});

async function loadClients() {
    try {
        const response = await fetch(`${API_BASE_URL}/clients`);
        const clients = await response.json();
        const list = document.getElementById('clients-list');
        if (list) {
            list.innerHTML = clients.map(client => `
                <div class="client">
                    <h3>${client.first_name} ${client.last_name}</h3>
                    <p>Email: ${client.email}</p>
                </div>
            `).join('');
        }
    } catch (error) {
        console.error('Error loading clients:', error);
    }
}