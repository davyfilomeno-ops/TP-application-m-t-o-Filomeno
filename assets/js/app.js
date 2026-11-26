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
    searchForm.addEventListener("submit", (event) => {
      const input = searchForm.querySelector('input[name="nom"]');
      const feedback = searchForm.querySelector("[data-search-feedback]");
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
