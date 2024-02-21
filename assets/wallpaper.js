document.addEventListener("DOMContentLoaded", function() {
    // Applique le fond d'écran sauvegardé au chargement de la page
    const savedWallpaper = localStorage.getItem("wallpaperPath");
    if (savedWallpaper) {
        document.body.style.backgroundImage = `url(${savedWallpaper})`;
    }

    // Ajoute des écouteurs d'événements aux images dans le containerWp
    const wallpapers = document.querySelectorAll(".containerWp .wp img");
    wallpapers.forEach(function(wallpaper) {
        wallpaper.addEventListener("click", function() {
            const imagePath = wallpaper.getAttribute("src");
            // Change le fond d'écran de la page
            document.body.style.backgroundImage = `url(${imagePath})`;
            // Sauvegarde le chemin de l'image dans localStorage
            localStorage.setItem("wallpaperPath", imagePath);
        });
    });
});
