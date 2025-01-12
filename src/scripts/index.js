import '../style/index.css';
const socket = new WebSocket("ws://localhost:8080/registration/");

const singInButton = document.querySelector('.login__sing-in');
const singUpButton = document.querySelector('.login__sing-up');
const popupAuthorization = document.querySelector('.popup_type_authorization');
const popupRegistration = document.querySelector('.popup_type_registration');

const registrationForm = document.forms.registration;
const registrationFormName = registrationForm.elements.name;
const registrationFormPassword = registrationForm.elements.password;
const authorizationForm = document.forms.authorization;

singInButton.addEventListener('click', function() {openPopupHandler(popupAuthorization);})
singUpButton.addEventListener('click', function() {openPopupHandler(popupRegistration);})

registrationForm.addEventListener('submit', (evt) => {
  evt.preventDefolt();
  handleRegistrationForm();
});
function handleRegistrationForm() {
  console.log("regist")
  if(username && password)
  {
    const message = {
      username: registrationFormName.value.trim(),
      password: registrationForm.value.trim(),
      email: ""
    }
    console.log(message)
    socket.send(JSON.stringify(message));
  }
  console.log("ok")
  
}

socket.onmessage = (evt) => {
  const receivedData = JSON.parse(evt.data);
  console.log(receivedData)
}

authorizationForm.addEventListener('submit', (evt) => {
  evt.preventDefolt();
  handleAuthorizationForm();
});
function handleAuthorizationForm() {

}

function openPopupHandler(popup) {
    openPopup(popup);
    document.addEventListener('keydown', closingOnEsc);
    popup.addEventListener('click', closingOnOverlay);
  }
  
  function closePopupHandler(popup) {
    removeClosingOnEsc();
    removeClosingOnOverlay(popup);
    closePopup(popup);
  }
  
  function removeClosingOnEsc() {
    document.removeEventListener('keydown', closingOnEsc);
  }
  
  function removeClosingOnOverlay(popup) {
    popup.removeEventListener('click', closingOnOverlay);
  }
  
  function closingOnEsc(evt) {
    if (evt.key === 'Escape') {
      const openPopupElement = document.querySelector('.popup_is-opened');
      if (openPopupElement) {
        closePopupHandler(openPopupElement);
      }
    }
  }
  
  function closingOnOverlay(evt) {
    if (evt.target === evt.currentTarget) {
      closePopupHandler(evt.currentTarget);
    }
  }
  
  function openPopup(popup) {
    popup.classList.add('popup_is-opened');
  }
  
  function closePopup(popup) {
    popup.classList.remove('popup_is-opened');
  }
  
//   export{openPopupHandler as open, closePopupHandler as close};


