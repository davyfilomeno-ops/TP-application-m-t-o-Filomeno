(() => {
  const weatherDescriptions = [
    { codes: [0], icon: "☀️", label: "Ciel dégagé" },
    { codes: [1], icon: "🌤️", label: "Majoritairement dégagé" },
    { codes: [2], icon: "⛅", label: "Partiellement nuageux" },
    { codes: [3], icon: "☁️", label: "Couvert" },
    { codes: [45, 48], icon: "🌫️", label: "Brouillard" },
    { codes: [51, 53, 55, 56, 57], icon: "🌦️", label: "Bruine" },
    { codes: [61, 63, 65, 66, 67], icon: "🌧️", label: "Pluie" },
    { codes: [71, 73, 75, 77], icon: "🌨️", label: "Neige" },
    { codes: [80, 81, 82], icon: "🌦️", label: "Averses" },
    { codes: [85, 86], icon: "🌨️", label: "Averses de neige" },
    { codes: [95, 96, 99], icon: "⛈️", label: "Orages" }
  ];

  document.querySelectorAll("[data-weather-code]").forEach((element) => {
    const code = Number(element.dataset.weatherCode);
    const description = weatherDescriptions.find((item) => item.codes.includes(code));

    if (!description) {
      if (element.matches("[data-weather-description]")) {
        element.textContent = `Code météo ${code}`;
      }
      return;
    }

    if (element.matches("[data-weather-symbol]")) {
      element.textContent = description.icon;
    }
    if (element.matches("[data-weather-description]")) {
      element.textContent = description.label;
    }
  });

  document.querySelectorAll("[data-forecast-date]").forEach((element) => {
    const date = new Date(`${element.textContent.trim()}T00:00:00Z`);
    if (!Number.isNaN(date.getTime())) {
      element.textContent = new Intl.DateTimeFormat("fr-FR", {
        weekday: "short",
        day: "numeric",
        month: "short",
        timeZone: "UTC"
      }).format(date);
    }
  });

  const resultCount = document.querySelector("[data-result-count]");
  if (resultCount) {
    const count = document.querySelectorAll(".commune-card").length;
    resultCount.textContent = `${count} ${count > 1 ? "résultats" : "résultat"}`;
  }

  const searchForm = document.querySelector("[data-search-form]");
  if (searchForm) {
    const input = searchForm.querySelector('input[name="nom"]');
    const feedback = searchForm.querySelector("[data-search-feedback]");
    const suggestionList = searchForm.querySelector("#commune-suggestions");
    let activeIndex = -1;
    let debounceTimer;
    let requestController;

    const normalizeName = (value) => value
      .normalize("NFD")
      .replace(/[\u0300-\u036f]/g, "")
      .toLocaleLowerCase("fr")
      .trim();

    const closeSuggestions = () => {
      suggestionList.hidden = true;
      input.setAttribute("aria-expanded", "false");
      input.removeAttribute("aria-activedescendant");
      activeIndex = -1;
    };

    const selectCommune = (commune) => {
      window.location.assign(`/communes/${encodeURIComponent(commune.code)}`);
    };

    const renderSuggestions = (communes) => {
      suggestionList.replaceChildren();
      activeIndex = -1;

      communes.forEach((commune, index) => {
        const item = document.createElement("li");
        item.setAttribute("role", "presentation");

        const option = document.createElement("button");
        option.id = `commune-suggestion-${index}`;
        option.className = "suggestion-option";
        option.type = "button";
        option.setAttribute("role", "option");
        option.setAttribute("aria-selected", "false");

        const name = document.createElement("span");
        name.className = "suggestion-name";
        name.textContent = commune.nom;

        const code = document.createElement("span");
        code.className = "suggestion-code";
        code.textContent = commune.code;

        option.append(name, code);
        option.addEventListener("click", () => selectCommune(commune));
        item.append(option);
        suggestionList.append(item);
      });

      if (communes.length === 0) {
        closeSuggestions();
        feedback.textContent = `Aucune commune ne commence par « ${input.value.trim()} ».`;
        return;
      }

      suggestionList.hidden = false;
      input.setAttribute("aria-expanded", "true");
      feedback.textContent = `${communes.length} communes trouvées. Choisissez-en une.`;
    };

    const loadSuggestions = async (prefix) => {
      if (requestController) {
        requestController.abort();
      }
      requestController = new AbortController();

      try {
        const response = await fetch(
          `/api/communes/suggestions?nom=${encodeURIComponent(prefix)}`,
          { signal: requestController.signal }
        );
        if (!response.ok) {
          throw new Error(`Suggestions indisponibles (${response.status}).`);
        }

        const communes = await response.json();
        if (normalizeName(input.value) !== normalizeName(prefix)) {
          return;
        }

        const normalizedPrefix = normalizeName(prefix);
        const matchingCommunes = communes.filter((commune) =>
          normalizeName(commune.nom).startsWith(normalizedPrefix)
        );
        renderSuggestions(matchingCommunes);
      } catch (error) {
        if (error.name === "AbortError") {
          return;
        }
        closeSuggestions();
        feedback.textContent = "Les suggestions sont indisponibles. Vous pouvez lancer la recherche.";
      }
    };

    input.addEventListener("input", () => {
      const prefix = input.value.trim();
      clearTimeout(debounceTimer);
      closeSuggestions();

      if (requestController) {
        requestController.abort();
      }

      if ([...prefix].length < 2) {
        feedback.textContent = prefix
          ? "Saisissez au moins deux lettres pour afficher les suggestions."
          : "Saisissez le nom d'une commune pour commencer.";
        return;
      }

      feedback.textContent = "Recherche des communes…";
      debounceTimer = setTimeout(() => loadSuggestions(prefix), 250);
    });

    input.addEventListener("keydown", (event) => {
      const options = [...suggestionList.querySelectorAll('[role="option"]')];
      if (suggestionList.hidden || options.length === 0) {
        if (event.key === "Escape") {
          closeSuggestions();
        }
        return;
      }

      if (event.key === "ArrowDown" || event.key === "ArrowUp") {
        event.preventDefault();
        const direction = event.key === "ArrowDown" ? 1 : -1;
        activeIndex = (activeIndex + direction + options.length) % options.length;

        options.forEach((option, index) => {
          const isActive = index === activeIndex;
          option.classList.toggle("is-active", isActive);
          option.setAttribute("aria-selected", String(isActive));
        });
        input.setAttribute("aria-activedescendant", options[activeIndex].id);
      } else if (event.key === "Enter" && activeIndex >= 0) {
        event.preventDefault();
        options[activeIndex].click();
      } else if (event.key === "Escape") {
        closeSuggestions();
      }
    });

    document.addEventListener("click", (event) => {
      if (!searchForm.contains(event.target)) {
        closeSuggestions();
      }
    });

    searchForm.addEventListener("submit", (event) => {
      const button = searchForm.querySelector('button[type="submit"]');
      const commune = input.value.trim();

      if (!commune) {
        event.preventDefault();
        input.focus();
        return;
      }

      feedback.textContent = `Recherche de « ${commune} »…`;
      button.disabled = true;
      button.querySelector("span").textContent = "Recherche…";
    });
  }
})();
