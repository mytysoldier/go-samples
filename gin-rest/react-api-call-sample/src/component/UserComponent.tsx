import { useEffect, useState } from "react";
import User from "../api/model/User";
import { fetchUserDatas } from "../api/UserApi";
import "./UserComponent.css";

const UserComponent: React.FC = () => {
  const [userList, setUserList] = useState<User[] | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchUserDatas();
      setUserList(data);
    };

    fetchData();
  }, []);

  return (
    <div>
      {userList ? (
        <div>
          <h2>User List</h2>
          <table className="user-table">
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Age</th>
              </tr>
            </thead>
            <tbody>
              {userList.map((user) => (
                <tr key={user.Id}>
                  <td>{user.Id}</td>
                  <td>{user.Name}</td>
                  <td>{user.Age}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <p>Loading user data...</p>
      )}
    </div>
  );
};

export default UserComponent;
