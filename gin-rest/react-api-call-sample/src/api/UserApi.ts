import User from "./model/User";

export const fetchUserData = async (id: number): Promise<User | null> => {
  try {
    // const response = await fetch("https://api.example.com/users");
    // const data: User[] = await response.json();
    // return data;

    // ダミーデータを生成して返す
    const dummyData: User = { Id: 1, Name: "John Doe", Age: 25 };

    return dummyData;
  } catch (error) {
    console.error("Error fetching user data:", error);
    return null;
  }
};

export const fetchUserDatas = async (): Promise<User[] | null> => {
  try {
    // const response = await fetch("https://api.example.com/users");
    // const data: User[] = await response.json();
    // return data;

    // ダミーデータを生成して返す
    const dummyData: User[] = [
      { Id: 1, Name: "John Doe", Age: 25 },
      { Id: 2, Name: "Jane Smith", Age: 30 },
      // 他にも必要なだけダミーデータを追加できます
    ];

    return dummyData;
  } catch (error) {
    console.error("Error fetching user data:", error);
    return null;
  }
};

export const UpdateUser = async (): Promise<User | null> => {};
