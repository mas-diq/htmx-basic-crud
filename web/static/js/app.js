// Initialize theme from localStorage or default to light
document.addEventListener('DOMContentLoaded', function () {
    const theme = localStorage.getItem('theme') || 'light';
    document.documentElement.setAttribute('data-theme', theme);
});

// Add HTMX enhancements
document.addEventListener('htmx:configRequest', function (event) {
    // For forms that should be submitted as application/x-www-form-urlencoded
    if (event.detail.elt.tagName === 'FORM' && !event.detail.elt.hasAttribute('enctype')) {
        event.detail.headers['Content-Type'] = 'application/x-www-form-urlencoded';
    }
});

// Add animations on HTMX events for smooth transitions
document.addEventListener('htmx:beforeSwap', function (event) {
    // Skip on page load
    if (!event.detail.requestConfig) return;

    // Add fade-out class before removing elements
    if (event.detail.target && event.detail.target.classList) {
        event.detail.target.classList.add('fade-out');
    }
});

document.addEventListener('htmx:afterSwap', function (event) {
    // Skip on page load
    if (!event.detail.requestConfig) return;

    // Add fade-in class to new elements
    if (event.detail.target) {
        // Remove fade-out if it was added
        event.detail.target.classList.remove('fade-out');

        // Add fade-in to all new child elements
        const newElements = event.detail.target.querySelectorAll('*');
        newElements.forEach(el => {
            el.classList.add('fade-in');

            // Remove the class after animation completes
            setTimeout(() => {
                el.classList.remove('fade-in');
            }, 500);
        });
    }
});

// Show toast notifications for successful operations
document.addEventListener('htmx:responseInfo', function (event) {
    const status = event.detail.xhr.status;
    const path = event.detail.pathInfo.requestPath;

    // Show success toast for successful CRUD operations
    if ((status === 200 || status === 204) &&
        (path.endsWith('/notes') || /\/notes\/\d+/.test(path))) {

        let message = '';

        if (event.detail.requestConfig.verb === 'post') {
            message = 'Note created successfully!';
        } else if (event.detail.requestConfig.verb === 'put') {
            message = 'Note updated successfully!';
        } else if (event.detail.requestConfig.verb === 'delete') {
            message = 'Note deleted successfully!';
        }

        if (message && window.showToast) {
            window.showToast(message);
        }
    }
});

// Simple toast notification function
window.showToast = function (message, type = 'success') {
    const toast = document.createElement('div');
    toast.className = `toast toast-${type} fade-in`;
    toast.textContent = message;

    document.body.appendChild(toast);

    setTimeout(() => {
        toast.classList.remove('fade-in');
        toast.classList.add('fade-out');

        setTimeout(() => {
            document.body.removeChild(toast);
        }, 500);
    }, 3000);
};