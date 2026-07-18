import { api } from "$lib/utils/api.svelte.js"

export const isLoggedIn = async () => {
	try {
		const resp = await api.post("/me");
	}
	catch (e) {
		console.error(e)
		return false
	}

	return true
}

export const normalizeDate = (date) => {
  return date ? date.split("T")[0] : null
}

export const handleImageChange = (file) => {
  return new Promise((resolve) => {
    if (!file) {
      resolve("");
      return;
    }
    
    const reader = new FileReader();
    
    reader.onload = (e) => {
      resolve(e.target.result);
    };
    
    reader.readAsDataURL(file);
  });
};

export const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  });
};
