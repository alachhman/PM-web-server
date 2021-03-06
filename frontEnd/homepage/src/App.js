import './App.css';
import Highlight from "react-highlight.js"

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h1>PokemasDB</h1>
                <p>
                    A REST API for anyone to pull data on a Pokemon Masters EX Character. <br/>
                    Not a Developer? <a href={"https://www.antnee.net"}>Click Here</a>
                </p>
                <img src={"/pokeball.png"} alt="logo" style={{height: "30px", width: "30px"}}/>
                <Highlight language={"javascript"} style={{textAlign: "left", width: "75vw", maxWidth:"900px"}}>
                    {
                        "const fetch = require(\"node-fetch\");\n\n" +
                        "const grabTrainer = async () => {\n" +
                        "     let trainerUri = \"http://pokemasdb.com/trainer\";\n" +
                        "     return await fetch(trainerUri)\n" +
                        "        .then(response => response.json())\n" +
                        "        .then(json => json.trainers)\n" +
                        "};\n\n" +
                        "grabTrainer().then(console.log);"
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