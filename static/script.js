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
    this.currentYear = "";
    this.loadYears();
  }

  async loadYears() {
    try {
      const response = await fetch("/api/years");
      const years = await response.json();
      this.renderYears(years);
      this.currentYear = years[0].name;
      this.loadSolutions(this.currentYear);
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
      card.className = "year-card";
      card.innerHTML = `
                <h3>${year.name}</h3>
            `;

      card.addEventListener("click", () => this.switchYear(year));
      this.yearSelector.appendChild(card);
    });
  }

  renderSolutions(year, solutions) {
    this.solutionsGrid.innerHTML = "";

    solutions.forEach((solution) => {
      const card = document.createElement("div");
      card.className = "solution-card";
      card.innerHTML = `
                <h3>${solution.name}</h3>
                <div class="status">${solution.status}</div>
            `;

      card.addEventListener("click", () =>
        this.runSolution(year, solution.day)
      );
      this.solutionsGrid.appendChild(card);
    });
  }

  async switchYear(year) {
    console.log("clicked", year.name);
    this.currentYear = year.name;
    this.loadSolutions(this.currentYear);
  }

  async runSolution(year, day) {
    this.outputElement.textContent = `Running ${day}...`;

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
      } else {
        this.outputElement.textContent =
          result.output || "No output (check if solution prints results)";
        this.solutionPart1Element.textContent = result.part1;
        this.solutionPart2Element.textContent = result.part2;
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
}

// Initialize the UI when page loads
document.addEventListener("DOMContentLoaded", () => {
  new AdventOfCodeUI();
});
