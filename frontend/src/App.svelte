<script>
  import { ParseInput } from "../wailsjs/go/main/App.js";
  import { ClickItemContainer } from "../wailsjs/go/main/App.js";

  let inputField = "";
  let results = [];

  // Debounce parseInput to avoid calling it on every keystroke immediately
  let debounceTimer;
  async function handleInput() {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(async () => {
      if (inputField.trim() === "") {
        results = [];
        return;
      }
      results = await ParseInput(inputField);
    }, 250); // Adjust debounce delay as needed (e.g., 250ms)
  }

  function handleClick(index) {
    ClickItemContainer(index);
    // Optionally clear input and results after click
    // inputField = "";
    // results = [];
  }

  // Example data for testing purposes if App.js is not available
  // setTimeout(() => {
  //   results = [
  //     { index: 0, title: "Visual Studio Code", subtitle: "Code editor by Microsoft", type: "Application" },
  //     { index: 1, title: "Terminal", subtitle: "Open a new terminal window", type: "Command" },
  //     { index: 2, title: "Firefox", subtitle: "Web browser by Mozilla", type: "Application" },
  //     { index: 3, title: "System Settings", subtitle: "Configure your system", type: "Application" },
  //     { index: 4, title: "Long Application Name That Might Overflow", subtitle: "A very descriptive subtitle for this application to test text overflow properties properly.", type: "Application" },
  //     { index: 5, title: "Calculator", subtitle: "Perform calculations", type: "Application" },
  //     { index: 6, title: "ls -la", subtitle: "List directory contents", type: "Command" },
  //     { index: 7, title: "Photos", subtitle: "View your image library", type: "Application" },
  //     { index: 8, title: "Notes", subtitle: "Jot down your thoughts", type: "Application" },
  //     { index: 9, title: "Another Item", subtitle: "Subtitle for another item", type: "Command"},
  //     { index: 10, title: "Yet Another App", subtitle: "A cool application", type: "Application"},
  //     { index: 11, title: "Final Command", subtitle: "The last command in this list", type: "Command"}
  //   ];
  // }, 500);
</script>

<main class="container">
  <div class="launcher-box">
    <div class="main-content-area">
      <div class="input-container-sticky">
        <div class="input-box">
          <input
            autocomplete="off"
            autocorrect="off"
            spellcheck="false"
            bind:value={inputField}
            class="input"
            id="InputField"
            type="text"
            on:input={handleInput}
            placeholder="Search for apps and commands..."
          />
        </div>
        <!-- <div class="top-right-buttons"></div> -->
      </div>

      <div class="scrollable-results-area">
        {#if results.length > 0}
          <ol class="results">
            {#each results as item (item.index)}
              <li
                class="result-item"
                on:click={() => handleClick(item.index)}
                tabindex="0"
                on:keydown={(e) => e.key === "Enter" && handleClick(item.index)}
              >
                <div class="result-icon">
                  <span class="icon-placeholder"
                    >{item.type === "Application" ? "🚀" : "❯_"}</span
                  >
                </div>
                <div class="result-text-content">
                  <span class="result-title" title={item.title}
                    >{item.title}</span
                  >
                  <span class="result-subtitle" title={item.subtitle}
                    >{item.subtitle}</span
                  >
                </div>
                {#if item.type === "Application"}
                  <span class="result-type app">App</span>
                {:else if item.type === "Command"}
                  <span class="result-type cmd">Cmd</span>
                {/if}
              </li>
            {/each}
          </ol>
        {:else if inputField.trim() !== ""}
          <div class="no-results">No results found for "{inputField}"</div>
        {/if}
      </div>
    </div>
  </div>
</main>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen,
      Ubuntu, Cantarell, "Open Sans", "Helvetica Neue", sans-serif;
    background-color: #1a1a1a; /* Even darker page background */
    display: flex;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    color: #f0f0f0; /* Lighter base text color for contrast */
  }

  .container {
    /* Container for the launcher */
  }

  .launcher-box {
    width: 750px;
    height: 450px;
    background-color: #282828; /* Darker launcher background */
    border-radius: 12px; /* Slightly more rounded corners */
    box-shadow: 0 8px 25px rgba(0, 0, 0, 0.35); /* Adjusted shadow for darker theme */
    display: flex;
    flex-direction: column;
    overflow: hidden;
    border: 1px solid #383838; /* Darker border */
  }

  .main-content-area {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    overflow: hidden;
  }

  .input-container-sticky {
    padding: 16px;
    background-color: #282828; /* Match darker launcher background */
    border-bottom: 1px solid #383838; /* Darker separator line */
  }

  .input {
    width: 100%;
    padding: 12px 15px;
    font-size: 1rem;
    background-color: #1e1e1e; /* Very dark input background */
    color: #f0f0f0; /* Light text color */
    border: 1px solid #484848; /* Darker border for input */
    border-radius: 8px; /* More rounded input field */
    box-sizing: border-box;
    outline: none;
    transition:
      border-color 0.2s ease,
      box-shadow 0.2s ease;
  }

  .input:focus {
    border-color: #007aff; /* Keep focus color for visibility */
    box-shadow: 0 0 0 3px rgba(0, 122, 255, 0.25); /* Slightly adjusted focus glow */
  }

  .input::placeholder {
    color: #777777; /* Adjusted placeholder color for dark theme */
  }

  .scrollable-results-area {
    flex-grow: 1;
    overflow-y: auto;
    overflow-x: hidden;
    padding: 0px 8px 8px 8px; /* Adjusted padding, no top padding for results list itself */
    background-color: #222222; /* Darker results background */

    /* Raycast-like scrollbar for WebKit browsers (Chrome, Safari, Edge) */
    scrollbar-width: thin; /* For Firefox */
    scrollbar-color: #555555 #222222; /* thumb track for Firefox */
  }

  .scrollable-results-area::-webkit-scrollbar {
    width: 6px; /* Thin scrollbar */
  }

  .scrollable-results-area::-webkit-scrollbar-track {
    background: #222222; /* Match results area background */
    border-radius: 3px;
  }

  .scrollable-results-area::-webkit-scrollbar-thumb {
    background-color: #555555; /* Scrollbar thumb color */
    border-radius: 3px;
    border: 1px solid #222222; /* Border to match track, creating an inset look */
  }

  .scrollable-results-area::-webkit-scrollbar-thumb:hover {
    background-color: #666666; /* Slightly lighter on hover */
  }

  .results {
    list-style-type: none;
    padding: 0;
    margin: 0;
  }

  .result-item {
    display: flex;
    align-items: center;
    padding: 10px 12px; /* Adjusted padding */
    border-bottom: 1px solid #333333; /* Darker item separator */
    cursor: pointer;
    transition: background-color 0.1s ease; /* Faster transition */
    border-radius: 6px;
    margin: 4px 0; /* Add some vertical margin between items */
  }

  .result-item:last-child {
    border-bottom: none;
  }

  .result-item:hover,
  .result-item:focus {
    background-color: #3a3a3a; /* Darker hover/focus state */
    outline: none;
  }

  .result-icon {
    margin-right: 14px; /* Slightly more space */
    font-size: 1.4rem;
    width: 24px;
    text-align: center;
    color: #c0c0c0; /* Lighter icon color */
  }

  .result-text-content {
    flex-grow: 1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    min-width: 0;
  }

  .result-title {
    font-weight: 500;
    color: #f0f0f0; /* Lighter title text */
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-right: 10px;
  }

  .result-subtitle {
    color: #909090; /* Adjusted subtitle color */
    font-size: 0.85rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    text-align: right;
    flex-shrink: 0;
  }

  .result-type {
    margin-left: 12px; /* Slightly more space */
    font-size: 0.75rem;
    color: #ffffff;
    padding: 3px 8px; /* Slightly more padding */
    border-radius: 5px; /* More rounded type tag */
    white-space: nowrap;
  }

  .result-type.app {
    background-color: #007aff; /* Keep distinct colors */
  }

  .result-type.cmd {
    background-color: #34c759;
  }

  .no-results {
    padding: 25px; /* More padding */
    text-align: center;
    color: #888888; /* Adjusted no-results text color */
    font-style: italic;
  }
</style>
