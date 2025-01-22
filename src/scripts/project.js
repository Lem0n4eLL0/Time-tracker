import '../style/project.css';
import Task from './components/task';
import Category from './components/category';
import {getTasks, addTask, getTaskCategories, deleteTask, changeTask} from './requests'
import * as taskCard from './components/taskCard'
import * as popup from './module';

const tasksInsertionPoint = document.querySelector('.tasks__list');
const projectID = window.location.href.match(/[1-9]+$/)[0];
const tasksList = new Map();
const categoriesList = [];

const addTaskButton = document.querySelector('.tasks__add-task');
const popapAddTask = document.querySelector('.popup_type_add-task')
const popapChangeTask = document.querySelector('.popup_type_change-task');
const popapDeleteTask = document.querySelector('.popup_type_delete');

const addTaskForm = document.forms.add_task;
const addTaskFormName = addTaskForm.elements.name;
const addTaskFormDescription = addTaskForm.elements.description;
const addTaskFormCategory = addTaskForm.elements.category;

const changeTaskForm = document.forms.change_task;
const changeTaskFormName = changeTaskForm.elements.name;
const changeTaskFormDescription = changeTaskForm.elements.description;
const changeTaskFormCategory = changeTaskForm.elements.category

const deleteTaskForm = document.forms.delete_task;

function openAddNewTaskPopap() {
  popup.open(popapAddTask);
}

function openChangeTaskPopap(element) {
  popup.open(popapChangeTask);
  changeTaskForm.dataset.taskid = element.dataset.taskid;
  setChangeFormField(Number(element.dataset.taskid));
}

function setChangeFormField(taskID) {
  const task = tasksList.get(taskID);
  changeTaskFormName.value = task.getTitle();
  changeTaskFormDescription.value = task.getDescription();
  for (let e of changeTaskFormCategory.options) {
    if(e.text === task.getCategory()) {
      changeTaskFormCategory.selectedIndex = e.index;
      break;
    }
  }
}

function openDeleteTaskPopap(element) {
  popup.open(popapDeleteTask);
  deleteTaskForm.dataset.taskid = element.dataset.taskid;
}

function showTasksHandler() {
  return getTasks(projectID)
  .then((res) => {
    if(res.ok) {
      return res.json();
    }
      return Promise.reject(res.status);
  })
  .then((res) => {
    if(res.Tasks != null) {
      res.Tasks.forEach(e => {
        tasksList.set(e.id, new Task(e.id, e.projectID, e.task_name, e.task_description, timeStructManger(e.end_date), timeStructManger(e.created_at), e.category_name, e.status));
        prependCard(tasksInsertionPoint, taskCard.create(tasksList.get(e.id), openDeleteTaskPopap, openChangeTaskPopap, statusTaskHandler));
      });
    }
  })
}

function getCategoriesHandler() {
  getTaskCategories()
  .then((res) => {
    if(res.ok) {
      return res.json()
    }
    return Promise.reject(res.status);
  })
  .then((res) => {
    fillCategoriesList(res);
    addCategoriesInDocument();
  })
  .catch(err => console.log(`Ошибка: ${err}`));
}

function fillCategoriesList(res) {
  res.forEach((e) => {
    categoriesList.push(new Category(e.id, e.name, e.description));
  });
}

function addCategoriesInDocument() {
  document.querySelectorAll('#category')
  .forEach((e) => {
    createCategoriesOption().forEach((category) => {
      e.appendChild(category);
    })
  })
}

function createCategoriesOption() {
  return categoriesList.map(e => e.createOption());
}

function addNewTaskHendler(evt) {
  evt.preventDefault();
  addTask({projectID: Number(projectID),
    task_name: addTaskFormName.value,
    task_description: addTaskFormDescription.value,
    category_name: addTaskFormCategory.options[addTaskFormCategory.selectedIndex].text})
  .then((res) => {
    if(res.ok) {
        return res.json();
    }
    return Promise.reject(res.status);
  })
  .then((res) => {
    tasksList.set(res.id, new Task(res.id, res.projectID, res.task_name, res.task_description, res.end_date, timeStructManger(res.created_at), res.category_name, res.status));
    prependCard(tasksInsertionPoint, taskCard.create(tasksList.get(res.id), openDeleteTaskPopap, openChangeTaskPopap, statusTaskHandler));

    popup.close(popapAddTask);
  })
  //.catch(err => console.log(`Ошибка: ${err}`));
}

function changeTaskHandler(evt) {
  evt.preventDefault();
  const taskID = Number(evt.target.dataset.taskid);
  const task = createChangeTaskfromFormValue(taskID);
  changeTask(task)
  .then((res) => {
    if(!res.ok) {
      return Promise.reject(res.status);
    }
    tasksList.set(taskID, task)
    taskCard.change(findTaskCard(taskID), task);
    popup.close(popapChangeTask);
  });
}

function createChangeTaskfromFormValue(taskID) {
  var task = new Task();
  tasksList.get(taskID).copyOf(task);
  task.setTitle(changeTaskFormName.value);
  task.setDescription(changeTaskFormDescription.value);
  task.setCategory(changeTaskFormCategory.options[changeTaskFormCategory.selectedIndex].text);
  return task;
}

function deleteTaskHandler(evt) {
  evt.preventDefault()
  const taskID = Number(evt.target.dataset.taskid)
  deleteTask(projectID, taskID)
  .then((res) => {
    if (!res.ok) {
      return Promise.reject(res.status);
    }
    tasksList.delete(taskID);
    taskCard.delete(findTaskCard(taskID));
    popup.close(popapDeleteTask);
  })
  .catch(err => console.log(`Ошибка: ${err}`));
}

function statusTaskHandler(evt) {
  evt.preventDefault();
  const taskID = Number(evt.target.closest('.task-card').dataset.taskid);
  var task;
  if(tasksList.get(taskID).getStatus()) {
    task = functionCreateUnexecutedTask(taskID);
  } else {
    task = functionCreateCompliteTask(taskID);
  }
  changeTask(task)
  .then((res) => {
    if(!res.ok) {
      return Promise.reject(res.status);
    }
    tasksList.set(taskID, task);
    taskCard.changeStatus(findTaskCard(taskID), task);
  })
}

function functionCreateCompliteTask(taskID) {
  var task = new Task();
  tasksList.get(taskID).copyOf(task);
  task.setEndDate(timeStructManger(new Date().toISOString()));
  console.log(task.getEndDate())
  task.setStatus(true);
  return task;
}

function timeStructManger(timeStr) {
  if(timeStr != null) {
    return timeStr.slice(0, 19).replace('T', ' ');
  }
  return null;
}

function functionCreateUnexecutedTask(taskID) {
  var task = new Task();
  tasksList.get(taskID).copyOf(task);
  task.setEndDate(null);
  task.setStatus(false);
  return task;
}

function findTaskCard(id) {
  for (let e of tasksInsertionPoint.querySelectorAll('.card')) {
    if(Number(e.dataset.taskid) === id) {
      return e;
    }
  }
}

function appendCard(insertionPoint, card)
{
  insertionPoint.append(card);
}

function prependCard(insertionPoint, card)
{
  insertionPoint.prepend(card);
}

function init() {
  showTasksHandler()
  .then(() => {
    getCategoriesHandler();
  })
  //.catch(err => console.log(`Ошибка: ${err}`));
}


addTaskButton.addEventListener('click', openAddNewTaskPopap);
addTaskForm.addEventListener('submit', addNewTaskHendler);
changeTaskForm.addEventListener('submit', changeTaskHandler);
deleteTaskForm.addEventListener('submit', deleteTaskHandler);


init()



