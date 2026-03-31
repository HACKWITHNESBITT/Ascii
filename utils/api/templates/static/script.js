// DOM Elements
const forms = document.querySelectorAll("form")
const loader = document.getElementById("loader")
const input = document.getElementById("textInput")
const asciiOutput = document.getElementById("ascii")
const logoOutput = document.getElementById("logo")

// Initialize animations on page load
document.addEventListener('DOMContentLoaded', () => {
    // Add fade-in animation to container
    const container = document.querySelector('.container')
    if (container) {
        container.classList.add('fade-in')
    }
    
    // Add focus animation to input
    if (input) {
        input.addEventListener('focus', () => {
            input.parentElement.classList.add('focused')
        })
        
        input.addEventListener('blur', () => {
            input.parentElement.classList.remove('focused')
        })
    }
    
    // Add hover effects to buttons
    const buttons = document.querySelectorAll('button')
    buttons.forEach(button => {
        button.addEventListener('mouseenter', () => {
            button.style.transform = 'translateY(-3px) scale(1.02)'
        })
        
        button.addEventListener('mouseleave', () => {
            button.style.transform = 'translateY(0) scale(1)'
        })
    })
})

// Form submission with loading animation
forms.forEach(form => {
    form.addEventListener("submit", (e) => {
        const text = input.value.trim()
        
        if (!text) {
            e.preventDefault()
            showNotification('Please enter some text', 'error')
            input.focus()
            return
        }
        
        if (form.id === "logoForm") {
            document.getElementById("logoText").value = text
        }
        
        if (form.id === "qrForm") {
            document.getElementById("qrText").value = text
        }
        
        // Show loader with animation
        loader.style.display = "block"
        loader.classList.add('fade-in')
        
        // Update download link
        document.getElementById("downloadASCII").href = "/download/ascii?text=" + encodeURIComponent(text)
        
        // Add loading state to button
        const button = form.querySelector('button')
        button.disabled = true
        button.innerHTML = '<span class="spinner"></span> Generating...'
    })
})

// Download QR function with animation
function copyASCII() {
    const ascii = asciiOutput.innerText.trim()
    
    if (!ascii) {
        showNotification('No QR code to download', 'error')
        return
    }
    
    // Create download link
    const blob = new Blob([ascii], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'qr-code.txt'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
    
    showNotification('QR code downloaded!', 'success')
    
    // Add success animation to output
    asciiOutput.classList.add('success')
    setTimeout(() => {
        asciiOutput.classList.remove('success')
    }, 500)
}

// Download Logo function with animation
function copyLogo() {
    const logo = logoOutput.innerText.trim()
    
    if (!logo) {
        showNotification('No logo to download', 'error')
        return
    }
    
    // Create download link
    const blob = new Blob([logo], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'logo.txt'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
    
    showNotification('Logo downloaded!', 'success')
    
    // Add success animation to output
    logoOutput.classList.add('success')
    setTimeout(() => {
        logoOutput.classList.remove('success')
    }, 500)
}

// Dark mode toggle with smooth transition
document.getElementById("darkToggle").onclick = () => {
    document.body.classList.toggle("dark")
    
    // Add transition effect
    document.body.style.transition = 'all 0.5s ease'
    
    // Save preference to localStorage
    const isDark = document.body.classList.contains("dark")
    localStorage.setItem('darkMode', isDark)
    
    // Show notification
    showNotification(isDark ? 'Dark mode enabled' : 'Light mode enabled', 'info')
}

// Load dark mode preference on page load
if (localStorage.getItem('darkMode') === 'true') {
    document.body.classList.add('dark')
}

// Notification system
function showNotification(message, type = 'info') {
    // Remove existing notification
    const existingNotification = document.querySelector('.notification')
    if (existingNotification) {
        existingNotification.remove()
    }
    
    // Create notification element
    const notification = document.createElement('div')
    notification.className = `notification notification-${type}`
    notification.innerHTML = `
        <span class="notification-message">${message}</span>
        <button class="notification-close" onclick="this.parentElement.remove()">×</button>
    `
    
    // Add styles
    notification.style.cssText = `
        position: fixed;
        top: 20px;
        left: 50%;
        transform: translateX(-50%);
        padding: 15px 25px;
        border-radius: 10px;
        background: ${type === 'success' ? '#22c55e' : type === 'error' ? '#ef4444' : '#3b82f6'};
        color: white;
        font-weight: 600;
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.3);
        z-index: 10000;
        display: flex;
        align-items: center;
        gap: 15px;
        animation: slideDown 0.3s ease-out;
    `
    
    // Add animation keyframes
    const style = document.createElement('style')
    style.textContent = `
        @keyframes slideDown {
            from {
                opacity: 0;
                transform: translateX(-50%) translateY(-20px);
            }
            to {
                opacity: 1;
                transform: translateX(-50%) translateY(0);
            }
        }
        
        .notification-close {
            background: none;
            border: none;
            color: white;
            font-size: 1.2rem;
            cursor: pointer;
            padding: 0;
            margin: 0;
            opacity: 0.8;
            transition: opacity 0.2s;
        }
        
        .notification-close:hover {
            opacity: 1;
        }
        
        .spinner {
            display: inline-block;
            width: 16px;
            height: 16px;
            border: 2px solid rgba(255, 255, 255, 0.3);
            border-radius: 50%;
            border-top-color: white;
            animation: spin 1s ease-in-out infinite;
        }
    `
    document.head.appendChild(style)
    
    // Add to page
    document.body.appendChild(notification)
    
    // Auto-remove after 3 seconds
    setTimeout(() => {
        if (notification.parentElement) {
            notification.style.animation = 'slideDown 0.3s ease-out reverse'
            setTimeout(() => notification.remove(), 300)
        }
    }, 3000)
}

// Add keyboard shortcuts
document.addEventListener('keydown', (e) => {
    // Ctrl/Cmd + Enter to generate QR
    if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        const qrForm = document.getElementById('qrForm')
        if (qrForm && input.value.trim()) {
            qrForm.submit()
        }
    }
    
    // Ctrl/Cmd + Shift + Enter to generate Logo
    if ((e.ctrlKey || e.metaKey) && e.shiftKey && e.key === 'Enter') {
        const logoForm = document.getElementById('logoForm')
        if (logoForm && input.value.trim()) {
            logoForm.submit()
        }
    }
    
    // Escape to clear input
    if (e.key === 'Escape') {
        input.value = ''
        input.focus()
    }
})

// Add input validation with visual feedback
input.addEventListener('input', () => {
    const text = input.value.trim()
    const isValid = text.length > 0
    
    if (isValid) {
        input.style.borderColor = '#22c55e'
    } else {
        input.style.borderColor = 'rgba(255, 255, 255, 0.2)'
    }
})

// Add tooltip functionality
function addTooltips() {
    const elements = document.querySelectorAll('[data-tooltip]')
    elements.forEach(element => {
        element.addEventListener('mouseenter', () => {
            const tooltip = document.createElement('div')
            tooltip.className = 'tooltip'
            tooltip.textContent = element.getAttribute('data-tooltip')
            tooltip.style.cssText = `
                position: absolute;
                bottom: 100%;
                left: 50%;
                transform: translateX(-50%);
                padding: 8px 12px;
                background: rgba(0, 0, 0, 0.8);
                color: white;
                font-size: 0.8rem;
                border-radius: 6px;
                white-space: nowrap;
                z-index: 1000;
                animation: fadeIn 0.2s ease-out;
            `
            element.style.position = 'relative'
            element.appendChild(tooltip)
        })
        
        element.addEventListener('mouseleave', () => {
            const tooltip = element.querySelector('.tooltip')
            if (tooltip) {
                tooltip.remove()
            }
        })
    })
}

// Initialize tooltips
addTooltips()

// Add smooth scrolling for anchor links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault()
        const target = document.querySelector(this.getAttribute('href'))
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth',
                block: 'start'
            })
        }
    })
})

// Add output animation when content changes
const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
        if (mutation.type === 'childList' && mutation.target.classList.contains('output')) {
            mutation.target.classList.add('fade-in')
            setTimeout(() => {
                mutation.target.classList.remove('fade-in')
            }, 500)
        }
    })
})

// Observe output elements for changes
if (asciiOutput) {
    observer.observe(asciiOutput, { childList: true })
}
if (logoOutput) {
    observer.observe(logoOutput, { childList: true })
}

// Add hover effect to outputs
if (asciiOutput) {
    asciiOutput.addEventListener('mouseenter', () => {
        asciiOutput.style.transform = 'scale(1.02)'
    })
    
    asciiOutput.addEventListener('mouseleave', () => {
        asciiOutput.style.transform = 'scale(1)'
    })
}

if (logoOutput) {
    logoOutput.addEventListener('mouseenter', () => {
        logoOutput.style.transform = 'scale(1.02)'
    })
    
    logoOutput.addEventListener('mouseleave', () => {
        logoOutput.style.transform = 'scale(1)'
    })
}

// Console welcome message
console.log('%c QR TEXTIFY ', 'background: linear-gradient(90deg, #22c55e, #3b82f6); color: white; font-size: 20px; font-weight: bold; padding: 10px; border-radius: 5px;')
console.log('%c Keyboard Shortcuts:', 'color: #22c55e; font-weight: bold;')
console.log('%c Ctrl/Cmd + Enter: Generate QR', 'color: #3b82f6;')
console.log('%c Ctrl/Cmd + Shift + Enter: Generate Logo', 'color: #eab308;')
console.log('%c Escape: Clear input', 'color: #8b5cf6;')
