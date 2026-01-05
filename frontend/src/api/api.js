const API_URL = "http://localhost:8080/api";

export const api = {
    getAll,
    getById,
    create,
    update,
    remove,
};



async function getAll(resource) {
    return request(`/${resource}`, "GET");
}

async function getById(resource, id) {
    return request(`/${resource}/${id}`, "GET");
}

async function create(resource, data) {
    return request(`/${resource}`, "POST", data);
}

async function update(resource, id, data) {
    return request(`/${resource}/${id}`, "PUT", data);
}

async function remove(resource, id) {
    return request(`/${resource}/${id}`, "DELETE");
}



async function request(url, method, body) {
    const options = {
        method,
        headers: {
            "Content-Type": "application/json",
        },
    };

    if (body !== undefined) {
        options.body = JSON.stringify(body);
    }

    const res = await fetch(API_URL + url, options);

    if (!res.ok) {
        const text = await res.text();
        throw new Error(text || "API error");
    }

    return res.status === 204 ? null : res.json();
}
