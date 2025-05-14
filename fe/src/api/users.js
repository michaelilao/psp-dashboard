import { api } from "./api";

async function fetchAllUsers() {
  const uri = new URL("user", api)
  try {
    const res = await fetch(uri)
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: []
    }
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: []
    }
  }
}


async function updateUserById(user) {
  const uri = new URL(`user/${user.id}`, api)
  try {
    const res = await fetch(uri, {
      method:"PUT",
      body: JSON.stringify(user)
    })
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: {}
    }
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: {}
    }
  }
}


async function createNewUser(user) {
  const uri = new URL(`user`, api)
  try {
    const res = await fetch(uri, {
      method:"POST",
      body: JSON.stringify(user)
    })
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: {}
    }
    
  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: {}
    }
  }
}

async function deleteUserById(user) {
  const uri = new URL(`user/${user.id}`, api)
  try {
    const res = await fetch(uri, {
      method:"DELETE",
    })
    if (res.ok) {
      const json = await res.json()
      return {
        error: false,
        data: json
      }    
    }
    const error = await res.json()
    const message = error?.error 
    return {
      error: true,
      message: message,
      data: {}
    }

  } catch(err) {
    console.error(err)
    return {
      error: true,
      data: {}
    }
  }
}


export {fetchAllUsers, updateUserById, deleteUserById, createNewUser}