import * as general from '../common';

const cardTamplate = document.querySelector('#card-project-template').content;

function createProjectCard(project, delEl, changeEl) {
  const cardElement = cardTamplate.querySelector('.card').cloneNode(true);
  cardElement.dataset.projectid = project.getID();
  cardElement.querySelector('.project-card__link').href = general.createHost(`projects/${project.getID()}`);
  cardElement.querySelector('.card__title').textContent = project.getTitle();
  cardElement.querySelector('.card__description').textContent = project.getDescription();
  cardElement.querySelector('.card__delete-button').addEventListener('click', () => delEl(cardElement));
  cardElement.querySelector('.card__change-button').addEventListener('click', () => changeEl(cardElement));
  return cardElement;
}

function delElement(element) {
  element.remove();
}

function changeElement(element, project) {
  element.querySelector('.card__title').textContent = project.getTitle();
  element.querySelector('.card__description').textContent = project.getDescription();
}

export {createProjectCard as create, delElement as delete, changeElement as change};

