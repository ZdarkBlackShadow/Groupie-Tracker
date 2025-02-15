document.addEventListener('DOMContentLoaded', function() {
    initializeFilters();
});

function initializeFilters() {
    // Sélectionner les éléments
    const openMenuButton = document.querySelector('#openMenuButton');
    const filterMenu = document.querySelector('#filterMenu');
    const overlay = document.querySelector('#overlay');
    const applyFiltersButton = document.querySelector('#applyFiltersButton');
    const resetFiltersButton = document.querySelector('#resetFiltersButton');

    // Vérifier si les éléments existent
    if (!openMenuButton || !filterMenu || !overlay) {
        console.error('Elements not found');
        return;
    }

    // Ouvrir le menu des filtres
    openMenuButton.addEventListener('click', () => {
        console.log('Opening filter menu'); // Debug
        filterMenu.style.display = 'block';
        overlay.style.display = 'block';
        document.body.style.overflow = 'hidden';
    });

    // Fermer le menu des filtres
    function closeFilterMenu() {
        console.log('Closing filter menu'); // Debug
        filterMenu.style.display = 'none';
        overlay.style.display = 'none';
        document.body.style.overflow = 'auto';
    }

    // Fermer sur overlay click
    overlay.addEventListener('click', closeFilterMenu);

    // Appliquer les filtres
    if (applyFiltersButton) {
        applyFiltersButton.addEventListener('click', () => {
            const formData = new FormData(document.getElementById('filterForm'));
            const url = new URL(window.location);

            // Nettoyer les anciens paramètres
            url.searchParams.delete('recipe_filter');
            url.searchParams.delete('types_filter');
            url.searchParams.delete('rarity_filter');
            url.searchParams.delete('sort_filter');

            // Ajouter les nouveaux paramètres
            const recipeFilter = formData.get('recipe_filter');
            if (recipeFilter) {
                url.searchParams.append('recipe_filter', recipeFilter);
            }

            formData.getAll('types_filter').forEach(type => {
                url.searchParams.append('types_filter', type);
            });

            formData.getAll('rarity_filter').forEach(rarity => {
                url.searchParams.append('rarity_filter', rarity);
            });

            const sortFilter = formData.get('sort_filter');
            if (sortFilter) {
                url.searchParams.append('sort_filter', sortFilter);
            }

            // Mettre à jour l'URL et recharger
            window.history.replaceState({}, '', url);
            closeFilterMenu();
            window.location.reload();
        });
    }

    // Réinitialiser les filtres
    if (resetFiltersButton) {
        resetFiltersButton.addEventListener('click', () => {
            const form = document.getElementById('filterForm');
            form.reset();
            
            const url = new URL(window.location);
            url.searchParams.delete('recipe_filter');
            url.searchParams.delete('types_filter');
            url.searchParams.delete('rarity_filter');
            url.searchParams.delete('sort_filter');
            
            window.history.replaceState({}, '', url);
            window.location.reload();
        });
    }

    // Initialiser les filtres depuis l'URL
    setFiltersFromURL();
}

function setFiltersFromURL() {
    const url = new URL(window.location);
    const form = document.getElementById('filterForm');
    
    if (!form) return;

    // Recipe filter
    const recipeFilter = url.searchParams.get('recipe_filter');
    if (recipeFilter) {
        const input = form.querySelector(`input[name="recipe_filter"][value="${recipeFilter}"]`);
        if (input) input.checked = true;
    }

    // Types filter
    url.searchParams.getAll('types_filter').forEach(type => {
        const input = form.querySelector(`input[name="types_filter"][value="${type}"]`);
        if (input) input.checked = true;
    });

    // Rarity filter
    url.searchParams.getAll('rarity_filter').forEach(rarity => {
        const input = form.querySelector(`input[name="rarity_filter"][value="${rarity}"]`);
        if (input) input.checked = true;
    });

    // Sort filter
    const sortFilter = url.searchParams.get('sort_filter');
    if (sortFilter) {
        const input = form.querySelector(`input[name="sort_filter"][value="${sortFilter}"]`);
        if (input) input.checked = true;
    }
}
