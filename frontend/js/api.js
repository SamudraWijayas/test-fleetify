const API_URL = "http://127.0.0.1:3000";

function apiRequest(method, endpoint, data = null) {
  if (typeof $ === "undefined") {
    throw new Error("jQuery belum diload");
  }

  const token = localStorage.getItem("token");

  const ajaxOptions = {
    url: API_URL + endpoint,
    type: method.toUpperCase(),
    dataType: "json",
    headers: {},
    contentType: "application/json; charset=utf-8",
  };

  if (token) {
    ajaxOptions.headers["Authorization"] = "Bearer " + token;
  }

  if (ajaxOptions.type === "GET" && data) {
    ajaxOptions.data = data;
  } 
  else if (data !== null) {
    ajaxOptions.data = JSON.stringify(data);
  }

  return $.ajax(ajaxOptions)
    .fail(function(xhr) {
      console.error("API Error:", {
        status: xhr.status,
        response: xhr.responseText,
      });
    });
}
