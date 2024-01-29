import { useEffect, useState } from "react";
import User from "../api/model/User";
import { fetchUserDatas } from "../api/UserApi";
import "./UserComponent.css";
import { Link } from "react-router-dom";

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
                  <td>
                    <Link to={`/user/update/${user.Id}`}>更新</Link>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      ) : (
        <p>ユーザーデータを読み込んでいます...</p>
      )}
    </div>
  );
};

export default UserComponent;
