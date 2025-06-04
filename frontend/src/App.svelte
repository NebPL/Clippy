<script>
  import { ParseInput } from "../wailsjs/go/main/App.js";
  import { ClickItemContainer } from "../wailsjs/go/main/App.js";

  let inputField;

  async function parseInput() {
    const ol = document.getElementById("ItemContainers");
    ol.innerHTML = "";

    const result = await ParseInput(inputField);

    for (let i = 0; i < result.length; i++) {
      addToOl(result[i].title, result[i].subtitle, result[i].index);
    }
  }

  function addToOl(title, subtitle, index) {
    const ol = document.getElementById("ItemContainers");
    const newLi = document.createElement("li");
    newLi.className = "result-item";

    const containerDiv = document.createElement("div");
    containerDiv.className = "result-content";

    const titleSpan = document.createElement("h3");
    titleSpan.className = "result-title";
    titleSpan.textContent = title;

    const subtitleSpan = document.createElement("p");
    subtitleSpan.className = "result-subtitle";
    subtitleSpan.textContent = subtitle;

    containerDiv.appendChild(titleSpan);
    containerDiv.appendChild(subtitleSpan);

    containerDiv.addEventListener("click", () => {
      ClickItemContainer(index);
    });

    newLi.appendChild(containerDiv);
    ol.appendChild(newLi);
  }
</script>

<main class="container">
  <div class="launcher-box">
    <div class="input-box" id="input">
      <input
        autocomplete="off"
        autocorrect="off"
        bind:value={inputField}
        class="input"
        id="InputField"
        type="text"
        on:input={parseInput}
        placeholder="Type to search..."
      />
    </div>

    <ol id="ItemContainers" class="results"></ol>
  </div>
</main>

<style>
  :root {
    font-family: Inter, Avenir, Helvetica, Arial, sans-serif;
    font-size: 16px;
    line-height: 1.5;
    font-weight: 400;
    color: white;
    background-color: transparent;

    font-synthesis: none;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
  }

  body {
    margin: 0;
    padding: 0;
    background-color: transparent;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    overflow: hidden;
  }

  /* Container mit fixierter Größe */
  .container {
    width: 750px;
    height: 450px;
    box-sizing: border-box;
    padding: 1rem;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  /* Launcher Box füllt Container aus */
  .launcher-box {
    width: 100%;
    height: 100%;
    background-color: #1a1a1a;
    border-radius: 16px;
    box-shadow: 0 0 40px rgba(0, 0, 0, 0.5);
    display: flex;
    flex-direction: column;
    gap: 0.8rem;
    padding: 1rem 1.5rem;
    box-sizing: border-box;
  }

  /* Input-Feld */
  input {
    width: 100%;
    border-radius: 8px;
    border: 1px solid #333;
    padding: 0.9em 1.2em;
    font-size: 1.2em;
    font-weight: 600;
    background-color: #121212;
    color: white;
    outline: none;
    box-sizing: border-box;
    transition: border-color 0.3s;
  }

  input::placeholder {
    color: #666;
  }

  input:focus {
    border-color: #555;
  }

  /* Ergebnisliste ohne Scrollbar und bündig nach links */
  .results {
    flex: 1;
    overflow-y: hidden; /* Scrollbar entfernen */
    display: flex;
    flex-direction: column;
    gap: 0.7rem;
    margin: 0;
    padding: 0;
    list-style: none;
    /* Folgende Scrollbar-Regeln sind jetzt redundant, können aber zur Sicherheit bleiben, falls overflow-y: auto später wieder aktiviert wird */
    scrollbar-width: none; /* Für Firefox */
  }

  .results::-webkit-scrollbar {
    display: none; /* Für Webkit-Browser (Chrome, Safari) */
  }

  /* Eintrag */
  .result-item {
    width: 100%;
    background-color: #121212;
    border-radius: 12px;
    padding: 1rem 1.2rem;
    color: white;
    display: flex;
    justify-content: flex-start; /* Inhalt des result-item nach links ausrichten */
    align-items: flex-start; /* Nach links oben ausrichten */
    cursor: pointer;
    box-sizing: border-box;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.8); /* dunklerer Schatten */
    border: 1px solid #222; /* dunklerer Rand */
    transition:
      background-color 0.2s ease,
      box-shadow 0.3s ease,
      border-color 0.3s ease;
  }

  .result-item:hover {
    background-color: #2b2b2b;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 1);
    border-color: #444;
  }

  .result-content {
    display: flex;
    flex-direction: column;
    justify-content: flex-start; /* Vertikale Ausrichtung oben */
    align-items: flex-start; /* Horizontale Ausrichtung links */
    text-align: left;
    width: 100%; /* Sicherstellen, dass es den verfügbaren Platz einnimmt */
  }

  /* Titel */
  .result-title {
    font-weight: 700;
    font-size: 1.2em;
    margin-bottom: 0.25rem;
    color: #f0f0f0;
  }

  /* Untertitel */
  .result-subtitle {
    font-size: 0.95em;
    color: #aaa;
    user-select: none;
  }

  /* Dark Mode fallback */
  @media (prefers-color-scheme: dark) {
    :root {
      color: white;
      background-color: transparent;
    }
  }
</style>
