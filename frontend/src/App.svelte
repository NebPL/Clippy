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

    // Container für die ganze Zeile (mit Rahmen und abgerundeten Ecken)
    const containerDiv = document.createElement("div");
    containerDiv.className = "result-content";

    // Title links
    const titleSpan = document.createElement("span");
    titleSpan.className = "result-title";
    titleSpan.textContent = title;

    // Subtitle rechts
    const subtitleSpan = document.createElement("span");
    subtitleSpan.className = "result-subtitle";
    subtitleSpan.textContent = subtitle;

    // Inhalt zusammenfügen
    containerDiv.appendChild(titleSpan);
    containerDiv.appendChild(subtitleSpan);

    containerDiv.addEventListener("click", () => {
      ClickItemContainer(index);
    });

    newLi.appendChild(containerDiv);
    ol.appendChild(newLi);
  }
</script>

<main>
  <div class="input-box" id="input">
    <input
      autocomplete="off"
      autocorrect="off"
      bind:value={inputField}
      class="input"
      id="InputField"
      type="text"
      on:input={parseInput}
    />
  </div>

  <ol id="ItemContainers"></ol>
</main>

<style>
  /* OL ohne Aufzählungszeichen und volle Breite */
  #ItemContainers {
    list-style: none;
    padding-left: 0;
    margin-top: 16px;
    width: 100%;
  }
  #ItemContainers li {
    margin-bottom: 8px;
    width: 100%;
  }

  /* Zeilen-Container mit Rahmen, abgerundeten Ecken und voller Breite */
  .result-content {
    display: flex;
    justify-content: space-between; /* Title links, Subtitle rechts */
    align-items: center;
    padding: 10px 14px;
    width: 100%; /* sorgt dafür, dass die Box die gesamte verfügbare Breite nutzt */
    border: 1px solid #ccc; /* sichtbare Umrandung */
    border-radius: 8px;
    cursor: pointer;
    transition:
      background-color 0.15s ease,
      box-shadow 0.15s ease;
  }
  .result-content:hover {
    background-color: #f5f5f5;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.08);
  }

  /* Title links, fett */
  .result-title {
    font-weight: 700;
    font-size: 1rem;
    color: #222;
    margin: 0;
  }

  /* Subtitle rechts, heller und kleiner */
  .result-subtitle {
    font-size: 0.9rem;
    color: #666;
    margin: 0;
    white-space: nowrap;
  }

  /* Input-Feld */
  .input-box .input {
    width: 100%;
    border: none;
    border-radius: 6px;
    outline: none;
    height: 36px;
    line-height: 36px;
    padding: 20 12px;
    background-color: #f7f7f7;
    font-size: 1rem;
    transition: background-color 0.2s ease;
  }
  .input-box .input:hover,
  .input-box .input:focus {
    background-color: #fff;
    box-shadow: 0 0 6px rgba(0, 0, 0, 0.1);
  }
</style>
