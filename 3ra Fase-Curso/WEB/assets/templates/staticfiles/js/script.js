// Función para simular salida de datos en la terminal
function logToTerminal(message) {
    const terminal = document.getElementById('terminal');
    const newLine = document.createElement('div');
    newLine.className = 'line';
    newLine.innerHTML = `<span class="prompt">></span> ${message}`;
    
    // Insertar antes del cursor
    terminal.insertBefore(newLine, terminal.lastElementChild);
}

// Ejemplo: podrías llamar esto cuando recibas un JSON de tu backend en Go
console.log("Usa logToTerminal('Hola desde Go') para probar.");