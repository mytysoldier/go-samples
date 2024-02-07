import User from "./model/User";

export const fetchUserData = async (id: number): Promise<User | null> => {
  try {
    const response = await fetch(`http://localhost:8080/user/${id}`);
    if (!response.ok) {
      console.error("Error fetching user data. Status:", response.status);
      return null;
    }

    const responseData: { user: User } = await response.json();
    const userData: User = responseData.user;
    return userData;
  } catch (error) {
    console.error("Error fetching user data:", error);
    return null;
  }
};

export const fetchUserDatas = async (): Promise<User[] | null> => {
  try {
    const response = await fetch("http://localhost:8080/users");
    if (!response.ok) {
      // レスポンスが成功でない場合の処理
      console.error("Error fetching user data. Status:", response.status);
      return [];
    }

    const { users }: { users: User[] } = await response.json();
    return users || [];
  } catch (error) {
    console.error("Error fetching user data:", error);
    return null;
  }
};

export const RegisterUser = async (userData: User): Promise<User | null> => {
  try {
    const response = await fetch("http://localhost:8080/user", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      console.error("Error registering user. Status:", response.status);
      return null;
    }

    const registeredUser: User = await response.json();
    return registeredUser;
  } catch (error) {
    console.error("Error registering user:", error);
    return null;
  }
};

export const UpdateUser = async (userData: User): Promise<User | null> => {
  try {
    const apiUrl = `http://localhost:8080/user/${userData.Id}`;

    const response = await fetch(apiUrl, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      console.error("Failed to update user");
      return null;
    }

    const updatedUser: User = await response.json();
    return updatedUser;
  } catch (error) {
    console.error("Error updating user:", error);
    return null;
  }
};

export const deleteUser = async (userId: number): Promise<boolean> => {
  try {
    const apiUrl = `http://localhost:8080/user/${userId}`;

    const response = await fetch(apiUrl, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      console.error("Failed to delete user");
      return false;
    }

    return true;
  } catch (error) {
    console.error("Error deleting user:", error);
    return false;
  }
};
