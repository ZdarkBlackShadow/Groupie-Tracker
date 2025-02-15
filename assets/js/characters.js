document.addEventListener('DOMContentLoaded', function() {
    initializeFilters();
    initializeCards();
});

function initializeFilters() {
    const openMenuButton = document.getElementById('openMenuButton');
    const filterMenu = document.getElementById('filterMenu');
    const overlay = document.getElementById('overlay');
    const applyFiltersButton = document.getElementById('applyFiltersButton');
    const resetFiltersButton = document.getElementById('resetFiltersButton');

    // Open filter menu
    openMenuButton?.addEventListener('click', () => {
        filterMenu.style.display = 'block';
        overlay.style.display = 'block';
        document.body.style.overflow = 'hidden';
    });

    // Close on overlay click
    overlay?.addEventListener('click', closeFilterMenu);

    // Apply filters
    applyFiltersButton?.addEventListener('click', () => {
        const formData = new FormData(document.getElementById('filterForm'));
        const filters = {
            gender: formData.getAll('gender_filter'),
            nations: formData.getAll('nations_filter'),
            rarity: formData.getAll('rarity_filter'),
            elements: formData.getAll('elements_filter'),
            sort: formData.get('sort_filter')
        };

        updateURLWithFilters(filters);
        closeFilterMenu();
        window.location.reload();
    });

    // Reset filters
    resetFiltersButton?.addEventListener('click', () => {
        document.getElementById('filterForm').reset();
        clearURLFilters();
    });

    // Initialize filters from URL
    setFiltersFromURL();
}

function initializeCards() {
    const cards = document.querySelectorAll('.character-card');
    
    cards.forEach(card => {
        // Animation d'entrée
        card.style.opacity = '0';
        card.style.transform = 'translateY(20px)';
        
        setTimeout(() => {
            card.style.opacity = '1';
            card.style.transform = 'translateY(0)';
        }, Array.from(cards).indexOf(card) * 100);

        // Effet de survol
        card.addEventListener('mouseenter', (e) => {
            const img = e.currentTarget.querySelector('img');
            const vision = e.currentTarget.querySelector('.character-vision');
            
            if (img) {
                img.style.transform = 'scale(1.1)';
            }
            if (vision) {
                vision.style.transform = 'translateY(-5px)';
            }
        });

        card.addEventListener('mouseleave', (e) => {
            const img = e.currentTarget.querySelector('img');
            const vision = e.currentTarget.querySelector('.character-vision');
            
            if (img) {
                img.style.transform = 'scale(1)';
            }
            if (vision) {
                vision.style.transform = 'translateY(0)';
            }
        });
    });
}

function closeFilterMenu() {
    const filterMenu = document.getElementById('filterMenu');
    const overlay = document.getElementById('overlay');
    
    filterMenu.style.display = 'none';
    overlay.style.display = 'none';
    document.body.style.overflow = 'auto';
}

function updateURLWithFilters(filters) {
    const url = new URL(window.location);
    
    // Clear existing parameters
    url.searchParams.delete('gender_filter');
    url.searchParams.delete('nations_filter');
    url.searchParams.delete('rarity_filter');
    url.searchParams.delete('elements_filter');
    url.searchParams.delete('sort_filter');

    // Add new parameters
    Object.entries(filters).forEach(([key, values]) => {
        if (Array.isArray(values)) {
            values.forEach(value => {
                url.searchParams.append(`${key}_filter`, value);
            });
        } else if (values) {
            url.searchParams.append(`${key}_filter`, values);
        }
    });

    window.history.replaceState({}, '', url);
}

function clearURLFilters() {
    const url = new URL(window.location);
    
    ['gender_filter', 'nations_filter', 'rarity_filter', 
     'elements_filter', 'sort_filter'].forEach(param => {
        url.searchParams.delete(param);
    });

    window.history.replaceState({}, '', url);
    window.location.reload();
}

function setFiltersFromURL() {
    const params = new URLSearchParams(window.location.search);
    const form = document.getElementById('filterForm');
    
    if (!form) return;

    // Reset all inputs
    form.reset();

    // Set checkboxes and radio buttons based on URL parameters
    params.forEach((value, key) => {
        const elements = form.querySelectorAll(`[name="${key}"]`);
        elements.forEach(element => {
            if (element.type === 'checkbox' || element.type === 'radio') {
                element.checked = element.value === value;
            }
        });
    });
}