import appConfig from "@/app.config.json"
import axios from "axios";

const POST_URL_PATH = appConfig.apiServer + "expenses";

export function createExpense(payload) {
  console.log("sending payload to server", JSON.stringify(payload));
  return new Promise((resolve, reject) => {
    axios.post(POST_URL_PATH, {
      expense: payload,
    })
      .then(response => resolve(response.data.expense))
      .catch(error => reject(error));
  });
}

export function listExpenses() {
  return new Promise((resolve, reject) => {
    axios.get(appConfig.apiServer + "expenses")
      .then(response => resolve(response.data.expenses))
      .catch(error => reject(error));
  });
}

export function deleteExpense (id) {
  return axios.delete(appConfig.apiServer + "expenses/" + id);
}