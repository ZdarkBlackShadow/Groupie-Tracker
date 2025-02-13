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
    const rarityFilter = formData.getAll('rarity_filter');
    const typeFilter = formData.getAll('type_filter');
    const attackFilter = formData.get('attack_filter');
    const sortFilter = formData.get('sort_filter');
    const url = new URL(window.location);

    // Supprimer les anciens filtres
    url.searchParams.delete('rarity_filter');
    url.searchParams.delete('type_filter');
    url.searchParams.delete('attack_filter');
    url.searchParams.delete('sort_filter');

    // Ajouter les nouveaux filtres
    rarityFilter.forEach(rarity_filter => {
        url.searchParams.append('rarity_filter', rarity_filter)
    });
    typeFilter.forEach(type_filter => {
        url.searchParams.append('type_filter', type_filter)
    });
    if (attackFilter) {
        url.searchParams.append('attack_filter', attackFilter);
    }
    if (sortFilter) {
        url.searchParams.append('sort_filter', sortFilter);
    }

    // Mettre à jour l'URL sans recharger la page
    window.history.replaceState({}, '', url);

    // Fermer le menu
    filterMenu.style.display = 'none';
    overlay.style.display = 'none';

    window.location.reload();
});

// Réinitialiser les filtres
resetFiltersButton.addEventListener('click', () => {
    const form = document.getElementById('filterForm');
    form.reset();

    const url = new URL(window.location);

    // Supprimer tous les filtres de l'URL
    url.searchParams.delete('filter');
    url.searchParams.delete('sort_filter');
    url.searchParams.delete('rarity_filter');

    // Mettre à jour l'URL sans recharger la page
    window.history.replaceState({}, '', url);

    console.log('Filtres réinitialisés');
}); 