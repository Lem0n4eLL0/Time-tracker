import '../style/profile.css';
import * as projectCards from './components/projectCard'
import * as popup from './module';
import Project from './components/project'
import {getProjects, addProject, deleteProject, changeProject, getPDF} from './requests'

const progectsInsertionPoint = document.querySelector('.projects__list');

const addProjectButton = document.querySelector('.projects__add-project');
const popupAddProject = document.querySelector('.popup_type_add-project');
const popapChangeProject = document.querySelector('.popup_type_change-project');
const popapDeleteProject = document.querySelector('.popup_type_delete');

const addProjectForm = document.forms.add_project;
const addProjectFormName = addProjectForm.elements.name;
const addProjectFormDescription = addProjectForm.elements.description;

const changeProjectForm = document.forms.change_project;
const changeProjectFormName = changeProjectForm.elements.name;
const changeProjectFormDescription = changeProjectForm.elements.description;

const deleteProjectForm = document.forms.delete_project;

const getPdfReportForm = document.forms.get_pdf_report_form;

const projectsList = new Map();

function openAddNewProjectPopap() {
  popup.open(popupAddProject);
}

function openChangeProjectPopap(element) {
  popup.open(popapChangeProject);
  changeProjectForm.dataset.projectid = element.dataset.projectid;
}

function openDeleteProjectPopap(element) {
  popup.open(popapDeleteProject);
  deleteProjectForm.dataset.projectid = element.dataset.projectid;
}

function addNewProjectHendler(evt) {
  evt.preventDefault();
  addProject({name: addProjectFormName.value, description: addProjectFormDescription.value})
  .then((res) => {
    if(res.ok) {
        return res.json();
    }
    return Promise.reject(res.status);
  })
  .then((res) => {
    projectsList.set(res.id, new Project(res.id, res.project_name, res.project_description));
    appendCard(progectsInsertionPoint, createProjectCard(projectsList.get(res.id)))
    addProjectForm.reset();
    popup.close(popupAddProject);
  })
  .catch(err => console.log(`Ошибка: ${err}`));
}

function showProjectHandler() {
  getProjects()
  .then((res) => {
    if(res.ok) {
      return res.json();
    }
      return Promise.reject(res.status);
  })
  .then((res) => {
    res.Projects.forEach(e => {
      projectsList.set(e.id, new Project(e.id, e.project_name, e.project_description));
      prependCard(progectsInsertionPoint, createProjectCard(projectsList.get(e.id)));
    });
  })
  .catch(err => console.log(`Ошибка: ${err}`));
}

function deleteProjectHandler(evt) {
  evt.preventDefault();
  const projectID = Number(evt.target.dataset.projectid);
  deleteProject(projectID)
    .then((res) => {
      if (res.ok) {
        return ;
      }
      return Promise.reject(res.status);
    })
    .then(() => {
      projectsList.delete(projectID);
      projectCards.delete(findProjectCard(projectID));
      popup.close(popapDeleteProject);
    })
    .catch(err => console.log(`Ошибка: ${err}`));
}

function changeProjectHandler(evt) {
  evt.preventDefault();
  const project = new Project(Number(evt.target.dataset.projectid), changeProjectFormName.value, changeProjectFormDescription.value);
  changeProject({id:project.getID(), name:project.getTitle(), description:project.getDescription()})
  .then((res) => {
    if (res.ok) {
      return ;
    }
    return Promise.reject(res.status);
    })
    .then(() => {
      projectsList.set(project.getID(), project)
      projectCards.change(findProjectCard(project.getID()), project);
      changeProjectForm.reset();
      popup.close(popapChangeProject);
    })
    .catch(err => console.log(`Ошибка: ${err}`));
}

function getReportPDF(evt) {
  evt.preventDefault();
  getPDF(null)
  .then((res) => {
    if(res.ok) {
      return response.blob();
    }
    return Promise.reject(res.status);
  })
  .then((blob) => {
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;
    link.download = 'report.pdf';
    console.log("Hehe");
    // link.click();
    // window.URL.revokeObjectURL(url);
  })
  //.catch(err => console.log(`Ошибка: ${err}`));
}

function findProjectCard(id) {
  for (let e of progectsInsertionPoint.querySelectorAll('.card')) {
    if(Number(e.dataset.projectid) === id) {
      return e;
    }
  }
}

function createProjectCard(project) {
  return projectCards.create(project, openDeleteProjectPopap, openChangeProjectPopap)
}

function appendCard(insertionPoint, card)
{
  insertionPoint.append(card);
}

function prependCard(insertionPoint, card)
{
  insertionPoint.prepend(card);
}


addProjectButton.addEventListener('click', openAddNewProjectPopap)
addProjectForm.addEventListener('submit', addNewProjectHendler)
changeProjectForm.addEventListener('submit', changeProjectHandler)
deleteProjectForm.addEventListener('submit', deleteProjectHandler)
getPdfReportForm.addEventListener('submit', getReportPDF)

showProjectHandler()

