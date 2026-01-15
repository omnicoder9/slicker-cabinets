// No imports, relying on global jQuery from CDN
// import $ from 'jquery'; 

$(document).ready(() => {
    // --- Theme Toggle ---
    const themeToggleBtn = $('#theme-toggle');
    const htmlElement = $('html');
    
    // Check local storage
    const savedTheme = localStorage.getItem('theme') || 'light';
    htmlElement.attr('data-bs-theme', savedTheme);
    
    themeToggleBtn.on('click', () => {
        const currentTheme = htmlElement.attr('data-bs-theme');
        const newTheme = currentTheme === 'light' ? 'dark' : 'light';
        
        htmlElement.attr('data-bs-theme', newTheme);
        localStorage.setItem('theme', newTheme);
    });

    // --- Nested Dropdown Mobile Handling ---
    $('.dropdown-menu .dropdown-toggle').on('click', function(e) {
        if (!$(this).next().hasClass('show')) {
            $(this).parents('.dropdown-menu').first().find('.show').removeClass('show');
        }
        
        const $subMenu = $(this).next('.dropdown-menu');
        if ($subMenu.length) {
            $subMenu.toggleClass('show');
            $(this).parent().toggleClass('show'); // Toggle dropend class if needed
            
            e.preventDefault();
            e.stopPropagation();
        }
    });

    // --- Contact Form Handling ---
    const contactForm = $('#contact-form');
    if (contactForm.length) {
        contactForm.on('submit', (e) => {
            e.preventDefault();
            
            const name = $('#name').val() as string;
            const email = $('#email').val() as string;
            const message = $('#message').val() as string;
            
            // Basic client-side validation
            if (!name || !email || !message) {
                showAlert('Please fill in all fields.', 'danger', '#alert-container');
                return;
            }

            // Secure/Sanitize input (handled by backend mostly, but we send JSON)
            const payload = {
                name: name,
                email: email,
                message: message
            };

            $.ajax({
                url: '/api/contact',
                method: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(payload),
                success: (response: any) => {
                    showAlert(response.message, 'success', '#alert-container');
                    (contactForm[0] as HTMLFormElement).reset();
                },
                error: (xhr) => {
                    const err = xhr.responseJSON ? xhr.responseJSON.error : 'An error occurred';
                    showAlert('Error: ' + err, 'danger', '#alert-container');
                }
            });
        });
    }

    // --- Login Form Handling ---
    const loginForm = $('#login-form');
    if (loginForm.length) {
        loginForm.on('submit', (e) => {
            e.preventDefault();
            
            const username = $('#username').val() as string;
            const password = $('#password').val() as string;
            
            const payload = {
                username: username,
                password: password
            };

            $.ajax({
                url: '/api/login',
                method: 'POST',
                contentType: 'application/json',
                data: JSON.stringify(payload),
                success: (response: any) => {
                    // In a real app, store token in localStorage/cookie
                    console.log('Token:', response.token);
                    showAlert(response.message, 'success', '#login-alert');
                    setTimeout(() => {
                        window.location.href = '/'; // Redirect to home/dashboard
                    }, 1000);
                },
                error: (xhr) => {
                     const err = xhr.responseJSON ? xhr.responseJSON.error : 'Invalid credentials';
                     showAlert(err, 'danger', '#login-alert');
                }
            });
        });
    }

    function showAlert(message: string, type: string, containerSelector: string) {
        const alertHtml = `
            <div class="alert alert-${type} alert-dismissible fade show" role="alert">
                ${message}
                <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
            </div>
        `;
        $(containerSelector).html(alertHtml);
    }
});
