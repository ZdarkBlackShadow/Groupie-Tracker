# Groupie-Tracker
{{define "elements"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/static/css/style.css">
    <link rel="stylesheet" href="/static/css/elements.css">
    <link rel="icon" href="/static/favicon.ico" type="image/x-icon">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>>Genshin</title>
</head>
<body>
    <header class="header">
        {{if eq .IsLogin true}}
            {{template "navbar2"}}
        {{else}}
            {{template "navbar1"}}
        {{end}}
    </header>
    <section aria-label="Newest Photos">
        <div class="carousel" data-carousel>
            <button class="carousel-button prev" data-carousel-button="prev">&#8656;</button>
            <button class="carousel-button next" data-carousel-button="next">&#8658;</button>
            <ul data-slides>
                {{range .Data}}
                    <li class="slide" data-active>
                        <a href="/elements/details?id={{.Id}}" class="center-link">
                            <div class="element-card">
                                <h1>Name: {{.Name}}</h1>
                                <img src="{{.ImageUrl}}" alt="{{.Name}}">
                            </div>
                        </a>
                    </li>
                {{end}}
            </ul>
        </div>
    </section>
    
    <script defer>
    document.addEventListener("DOMContentLoaded", () => {
  const buttons = document.querySelectorAll("[data-carousel-button]");

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const offset = button.dataset.carouselButton === "next" ? 1 : -1;
      const slides = button.closest("[data-carousel]").querySelector("[data-slides]");

      const activeSlide = slides.querySelector("[data-active]");
      let newIndex = [...slides.children].indexOf(activeSlide) + offset;

      if (newIndex < 0) newIndex = slides.children.length - 1;
      if (newIndex >= slides.children.length) newIndex = 0;

      slides.children[newIndex].dataset.active = true;
      delete activeSlide.dataset.active;
    });
  });
});
    </script>
</body>
</html>
{{end}}