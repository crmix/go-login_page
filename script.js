const container = document.getElementById('container');
const registerBtn = document.getElementById('register');
const loginBtn = document.getElementById('login');

registerBtn.addEventListener('click', () => {
    container.classList.add("active");
});

loginBtn.addEventListener('click', () => {
    container.classList.remove("active");
});

// Function to handle sign-up
function signUp() {

    const name = document.getElementById('nameInput').value;
    const email = document.getElementById('emailInput').value;
    const password = document.getElementById('passwordInput').value;

     console.log(name, email, password)
    // Make a POST request to the register Api endpoint
    fetch('http://localhost:8004/auth/sign-up', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ name, email, password })
    })
    .then(response => {
        if (response.ok) {
            showPopupMessage('Register successful');
            setTimeout(hidePopupMessage, 3000); 
        } else {
            showPopupMessage('User exists with this email :(');
            setTimeout(hidePopupMessage, 3000);
        }
    })
    .catch(error => {
        console.error('Error signing up:', error);
    });
}

// Function to handle sign-in
function signIn() {
    
    const email = document.getElementById('emailInput').value; 
    const password = document.getElementById('passwordInput').value;
    
    console.log(email, password)
    // Make a POST request to the login Api endpoint
    fetch('http://localhost:8004/auth/sign-in', { 
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password }) 
    })
    .then(response => {
        if (response.ok) {
            showPopupMessage('Login successful');
            setTimeout(hidePopupMessage, 3000); 
        } else {
            showPopupMessage('Email or password is incorrect');
            setTimeout(hidePopupMessage, 3000);
        }
    })
    .catch(error => {
        console.error('Error signing up:', error);
    });
}

const signUpForm = document.querySelector('.sign-up form');
const signInForm = document.querySelector('.sign-in form');

signUpForm.addEventListener('submit', function(event) {
    event.preventDefault();
    
    signUp();
});


signInForm.addEventListener('submit', function(event) {
    event.preventDefault(); 
    signIn();
});

function showPopupMessage(message) {
    const popup = document.createElement('div');
    popup.classList.add('popup');
    popup.textContent = message;
    document.body.appendChild(popup);
}

function hidePopupMessage() {
    const popup = document.querySelector('.popup');
    if (popup) {
        popup.remove();
    }
}