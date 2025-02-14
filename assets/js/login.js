/*=============== SHOW MENU ===============*/
const navMenu = document.getElementById('nav-menu'),
      navToggle = document.getElementById('nav-toggle'),
      navClose = document.getElementById('nav-close')

/* Menu show */
if(navToggle){
   navToggle.addEventListener('click', () =>{
      navMenu.classList.add('show-menu')
   })
}

/* Menu hidden */
if(navClose){
   navClose.addEventListener('click', () =>{
      navMenu.classList.remove('show-menu')
   })
}

/* Login/Register Panel Switch */
document.addEventListener('DOMContentLoaded', () => {
    const container = document.getElementById("container");
    const overlayBtn = document.getElementById("overlayBtn");
    
    if (overlayBtn) {
        overlayBtn.addEventListener('click', () => {
            container.classList.toggle('right-panel-active');
        
            overlayBtn.classList.remove('btnScaled');
            window.requestAnimationFrame(() => {
                overlayBtn.classList.add('btnScaled');
            });
        });
    }
});