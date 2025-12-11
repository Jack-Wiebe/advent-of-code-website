class AdventOfCodeUI {
  constructor() {
    this.solutionsGrid = document.getElementById("solutionsGrid");
    this.yearSelector = document.getElementById("yearSelector");
    this.outputElement = document.getElementById("console-output");
    this.solutionPart1Element = document.getElementById(
      "part1-solution-output"
    );
    this.solutionPart2Element = document.getElementById(
      "part2-solution-output"
    );
    this.soltionDelta = document.getElementById("solution-delta");
    this.currentYearElement = "";
    this.currentSolutionElement = "";
    this.loadYears();
  }

  async loadYears() {
    try {
      const response = await fetch("/api/years");
      const years = await response.json();
      this.renderYears(years);
      this.loadSolutions(years[0].name);
    } catch (error) {
      console.error("Error loading years:", error);
      this.outputElement.textContent = "Error loading years";
    }
  }

  async loadSolutions(year) {
    try {
      const response = await fetch("/api/solutions/" + year);
      const solutions = await response.json();
      this.renderSolutions(year, solutions);
    } catch (error) {
      console.error("Error loading solutions:", error);
      this.outputElement.textContent = "Error loading solutions";
    }
  }

  renderYears(years) {
    this.yearSelector.innerHTML = "";

    years.forEach((year) => {
      const card = document.createElement("div");

      if (!this.currentYearElement) {
        this.currentYearElement = card;
        card.className = "year-card active";
      } else {
        card.className = "year-card";
      }

      card.innerHTML = `
                <h3>${year.name}</h3>
            `;

      card.addEventListener("click", (e) =>
        this.switchYear(year, e.currentTarget)
      );
      this.yearSelector.appendChild(card);
    });
  }

  renderSolutions(year, solutions) {
    this.solutionsGrid.innerHTML = "";

    solutions.sort((a, b) => {
      let day1 = a.day.split("-")[1];
      let day2 = b.day.split("-")[1];
      return day1 - day2;
    });

    solutions.forEach((solution) => {
      const card = document.createElement("div");
      card.className = "solution-card";
      card.innerHTML = `
                <h3>${solution.name}</h3>
                <div class="status">${solution.status}</div>
            `;

      card.addEventListener("click", (e) =>
        this.runSolution(year, solution.day, e.currentTarget)
      );
      this.solutionsGrid.appendChild(card);
    });
  }

  async switchYear(year, card) {
    console.log("clicked", year.name);

    if (this.currentYearElement) {
      this.currentYearElement.classList.remove("active");
    }
    this.currentYearElement = card;
    card.classList.add("active");

    this.loadSolutions(year.name);
  }

  async runSolution(year, day, card) {
    this.outputElement.textContent = `Running ${day}...`;

    if (this.currentSolutionElement) {
      this.currentSolutionElement.classList.remove("active");
    }
    this.currentSolutionElement = card;
    card.classList.add("active");

    // Add loading state to all cards
    document.querySelectorAll(".solution-card").forEach((card) => {
      card.classList.add("loading");
    });

    try {
      const response = await fetch(`/api/run/${year}/${day}`, {
        method: "POST",
      });
      const result = await response.json();

      if (result.error) {
        this.outputElement.textContent = `Error: ${result.error}\n\nOutput:\n${result.output}`;
        this.solutionPart1Element.textContent = "";
        this.solutionPart2Element.textContent = "";
        this.soltionDelta.textContent = "";
      } else {
        this.solutionPart1Element.textContent = result.part1;
        this.solutionPart2Element.textContent = result.part2;
        this.soltionDelta.textContent = result.delta + "ms";
      }
      if (result.output) {
        this.outputElement.textContent =
          result.output || "No output (check if solution prints results)";
      }
    } catch (error) {
      this.outputElement.textContent = `Network error: ${error.message}`;
    } finally {
      // Remove loading state
      document.querySelectorAll(".solution-card").forEach((card) => {
        card.classList.remove("loading");
      });
    }
  }

  enterFullscreen() {
    if (this.outputElement.requestFullscreen) {
      this.outputElement.requestFullscreen();
    } else if (this.outputElement.webkitRequestFullscreen) {
      this.outputElement.webkitRequestFullscreen();
    } else if (this.outputElement.msRequestFullscreen) {
      this.outputElement.msRequestFullscreen();
    }

    this.outputElement.classList.add("fullscreen-active");
  }

  exitFullscreen() {
    if (document.exitFullscreen) {
      document.exitFullscreen();
    } else if (document.webkitExitFullscreen) {
      document.webkitExitFullscreen();
    } else if (document.msExitFullscreen) {
      document.msExitFullscreen();
    }

    this.outputElementclassList.remove("fullscreen-active");
  }

  isFullscreen() {
    return (
      document.fullscreenElement ||
      document.webkitFullscreenElement ||
      document.msFullscreenElement
    );
  }
}

let uiInstance = null;
document.addEventListener("DOMContentLoaded", () => {
  uiInstance = new AdventOfCodeUI();
});

document.addEventListener("keydown", (e) => {
  if (e.key === "f" || e.key === "F") {
    if (uiInstance.isFullscreen()) {
      uiInstance.exitFullscreen();
    } else {
      uiInstance.enterFullscreen();
    }
  }

  if (e.key === "Escape" && uiInstance.isFullscreen()) {
    uiInstance.exitFullscreen();
  }
});
