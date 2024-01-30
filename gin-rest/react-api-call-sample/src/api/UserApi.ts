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

    // // ダミーデータを生成して返す
    // const dummyData: User = { Id: 1, Name: "John Doe", Age: 25 };

    // return dummyData;
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

    // // ダミーデータを生成して返す
    // const dummyData: User[] = [
    //   { Id: 1, Name: "John Doe", Age: 25 },
    //   { Id: 2, Name: "Jane Smith", Age: 30 },
    //   // 他にも必要なだけダミーデータを追加できます
    // ];

    // return dummyData;
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
