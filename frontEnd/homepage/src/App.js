import './App.css';
import Highlight from "react-highlight.js"

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h1>PokemasDB</h1>
                <p>
                    A REST API for anyone to pull data on a Pokemon Masters EX Character.
                </p>
                <img src={"/pokeball.png"} alt="logo" style={{height: "30px", width: "30px"}}/>
                <Highlight language={"javascript"} style={{textAlign: "left"}}>
                    {
                        "const grabTrainer = async () => {\n" +
                        "     return await fetch(\"http://pokemasdb.com/trainer\")\n" +
                        "        .then(response => response.json())\n" +
                        "        .then(json => json.trainers)\n" +
                        "}"
                    }
                </Highlight>
                <img src={"/pokeball.png"} alt="logo" style={{height: "30px", width: "30px"}}/>
                <p>
                    Documentation on <a href={"https://github.com/alachhman/PM-web-server"}>Github</a> | <a href={"https://discord.gg/XEyTDE9"}>Discord Support Server</a>
                </p>
            </header>
        </div>
    );
}

export default App;

const grabTrainer = async () => {
     return await fetch("http://pokemasdb.com/trainer")
        .then(response => response.json())
        .then(json => json.trainers)
}