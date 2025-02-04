import './style/index.css';
import * as popup from './scripts/module';
import * as requests from './scripts/requests'
import {enableValidation} from './scripts/validation'

const singInButton = document.querySelector('.header__log-in-button');
const singUpButton = document.querySelector('.header__sign-up-button');
const popupAuthorization = document.querySelector('.popup_type_authorization');
const popupRegistration = document.querySelector('.popup_type_registration');

const registrationForm = document.forms.registration;
const registrationFormName = registrationForm.elements.name;
const registrationFormPassword = registrationForm.elements.password;
const authorizationForm = document.forms.authorization;
const authorizationFormName = authorizationForm.elements.name;
const authorizationFormPassword = authorizationForm.elements.password;


//регестрация
registrationForm.addEventListener('submit', (evt) => {
  evt.preventDefault();
  handleRegistrationForm();
});

function handleRegistrationForm() {
  if(registrationFormName.value && registrationFormPassword.value)
  {
    requests.registration({
      username: registrationFormName.value.trim(),
      password: registrationFormPassword.value.trim(),
      email: null
    })
  }
}

// авторизация
authorizationForm.addEventListener('submit', (evt) => {
  evt.preventDefault();
  handleLoginForm();
});

function handleLoginForm() {
  if(authorizationFormName.value && authorizationFormPassword.value)
  {
    requests.login({
      username: authorizationFormName.value.trim(),
      password: authorizationFormPassword.value.trim(),
    })
  }
}

//Убрать когда пользователь авторизован!
singInButton.addEventListener('click', function() {popup.open(popupAuthorization);});
singUpButton.addEventListener('click', function() {popup.open(popupRegistration);});
enableValidation();



