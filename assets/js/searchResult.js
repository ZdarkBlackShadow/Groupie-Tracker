const openMenuButton = document.getElementById('openMenuButton');
const filterMenu = document.getElementById('filterMenu');
const overlay = document.getElementById('overlay');
const applyFiltersButton = document.getElementById('applyFiltersButton');
const resetFiltersButton = document.getElementById('resetFiltersButton');

// Ouvrir le menu des filtres
openMenuButton.addEventListener('click', () => {
    filterMenu.style.display = 'block';
    overlay.style.display = 'block';
});

// Fermer le menu des filtres lorsqu'on clique en dehors
overlay.addEventListener('click', () => {
    filterMenu.style.display = 'none';
    overlay.style.display = 'none';
});

// Appliquer les filtres et mettre à jour l'URL
applyFiltersButton.addEventListener('click', () => {
    const formData = new FormData(document.getElementById('filterForm'));
    const sortFilter = formData.getAll('sort_filter');
    const url = new URL(window.location);
    
    url.searchParams.delete('sort_filter');
    
    sortFilter.forEach(sort_filter => {
        url.searchParams.append('sort_filter', sort_filter)
    });
    
    window.history.replaceState({}, '', url);
    filterMenu.style.display = 'none';
    overlay.style.display = 'none';
    window.location.reload();
});

// Réinitialiser les filtres
resetFiltersButton.addEventListener('click', () => {
    const form = document.getElementById('filterForm');
    form.reset();

    const url = new URL(window.location);
    url.searchParams.delete('sort_filter');

    window.history.replaceState({}, '', url);
}); 