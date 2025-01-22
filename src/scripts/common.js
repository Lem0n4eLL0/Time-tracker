import '../style/variables.css';
import '../style/globals.css';
import '../style/general.css';

const host = "http://localhost:8080/";

function createHost(path) {
    return host + path;
}

export {createHost}
