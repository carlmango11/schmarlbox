import {useEffect, useState} from 'react'
import "./wasm_exec.js";

const FPS = 10;
const HEIGHT = 25;
const WIDTH = 80;

import './App.css';

function initWasm(setDisplay: (d: string[][]) => void) {
    const go = new window.Go();

    WebAssembly.instantiateStreaming(fetch("box.wasm"), go.importObject).then(
        (result) => {
            console.log("starting wasm");
            go.run(result.instance);
        }
    );

    startDisplay(setDisplay);
}

function startDisplay(setDisplay: (d: string[][]) => void) {
    setInterval(() => setDisplay(getLatestDisplay()), 1000 / FPS);
}

function getLatestDisplay(): string[][] {
    const state = window.getDisplay();

    const display: string[][] = [];

    for (let y = 0; y < HEIGHT; y++) {
        const row: string[] = [];

        for (let x = 0; x < WIDTH; x++) {
            const i = x + (y * WIDTH);
            const v: string = state[i];

            row.push(v);
        }

        display.push(row);
    }

    return display;
}

function App() {
    const [display, setDisplay] = useState<string[][]>([[]]);

    useEffect(() => {
        console.log("init wasm");
        initWasm(setDisplay);
    }, []);

    return <div className="display">
        { display.map(row => row.join("")).join("\n")}
    </div>
}

export default App
