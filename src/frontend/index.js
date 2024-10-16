const { app, BrowserWindow } = require('electron');
const path = require('path');

app.commandLine.appendSwitch('disable-gpu');
app.commandLine.appendSwitch('disable-software-rasterizer');
app.commandLine.appendSwitch('disable-renderer-backgrounding');

function createWindow() {
    const win = new BrowserWindow({
        width: 900,
        height: 890,
        // frame: false,
        webPreferences: {
            preload: path.join(__dirname, 'preload.js'),
            nodeIntegration: true
        }
    });

    win.loadFile('index.html');

    win.on('closed', () => {
        app.quit();
    });
}

app.whenReady().then(() => {
    createWindow();

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) {
            createWindow();
        }
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') {
        app.quit();
    }
});
