import * as general from '../common';

const cardTamplate = document.querySelector('#card-task-template').content;

function createTaskCard(task, delEl, changeEl, statusTaskHandler) {
  const cardElement = cardTamplate.querySelector('.task-card').cloneNode(true);
  cardElement.dataset.taskid = task.getID();

  cardElement.querySelector('.card__title').textContent = task.getTitle();
  cardElement.querySelector('.card__description').textContent = task.getDescription();
  const status = cardElement.querySelector('.task-card__status');
  status.checked = task.getStatus();
  const statrtDate = cardElement.querySelector('.task-card__start-date');
  statrtDate.textContent = task.getCreatedAt();
  statrtDate.setAttribute('datetime', task.getCreatedAt());
  const endDate = cardElement.querySelector('.task-card__end-date');
  endDate.textContent = task.getEndDate();
  endDate.setAttribute('datetime', task.getEndDate());
  if(task.getCategory() != "NONE") {
    cardElement.querySelector('.task-card__category').textContent = task.getCategory();
  }

  const submitButton = cardElement.querySelector('.complite-task-form__button');
  if(task.getStatus()) {
    cardElement.classList.add('task-card_complite');
    submitButton.textContent = "complite";
  } else {
    cardElement.classList.add('task-card_unexecuted');
    submitButton.textContent = "unexecuted";
  }

  cardElement.querySelector('.complite-task-form').addEventListener('submit', statusTaskHandler);
  cardElement.querySelector('.card__delete-button').addEventListener('click', () => delEl(cardElement));
  cardElement.querySelector('.card__change-button').addEventListener('click', () => changeEl(cardElement));
  return cardElement;
}

function delElement(element) {
  element.remove();
}

function changeElement(element, task) {
  element.querySelector('.card__title').textContent = task.getTitle();
  element.querySelector('.card__description').textContent = task.getDescription();
  element.querySelector('.card__category').textContent = task.getCategory();
}

function changeStatus(element, task) {
  element.querySelector('.card__end-date').textContent = task.getEndDate();
  const status = element.querySelector('.card__status');
  const submitButton = element.querySelector('.complite-task-form__button');
  status.checked = task.getStatus();
  if (status.checked) {
    element.classList.add('task-card_complite');
    element.classList.remove('task-card_unexecuted');
    submitButton.textContent = "complite";
    // submitButton.classList.add("complite-task-form__button_animation");
  } else {
    element.classList.add('task-card_unexecuted');
    element.classList.remove('task-card_complite');
    submitButton.textContent = "unexecuted";

  }
  // потом сделать множество вариаций
}

export {createTaskCard as create, delElement as delete, changeElement as change, changeStatus};
