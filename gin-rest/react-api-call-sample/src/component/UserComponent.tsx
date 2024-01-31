import { useEffect, useState } from "react";
import User from "../api/model/User";
import { deleteUser, fetchUserDatas } from "../api/UserApi";
import "./UserComponent.css";
import { Link } from "react-router-dom";

const UserComponent: React.FC = () => {
  const [userList, setUserList] = useState<User[] | null>(null);
  const [isDeleteDialogOpen, setDeleteDialogOpen] = useState<boolean>(false);
  const [deletingUserId, setDeletingUserId] = useState<number | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchUserDatas();
      setUserList(data);
    };

    fetchData();
  }, []);

  const handleDeleteClick = (userId: number) => {
    setDeletingUserId(userId);
    setDeleteDialogOpen(true);
  };

  const handleDeleteConfirm = async () => {
    if (deletingUserId !== null) {
      try {
        const isDeleted = await deleteUser(deletingUserId);

        if (isDeleted) {
          setDeleteDialogOpen(false);
          const updatedData = await fetchUserDatas();
          setUserList(updatedData);
        } else {
          console.error("Error deleting user");
        }
      } catch (error) {
        console.error("Error deleting user data:", error);
      }
    }
  };

  const handleDeleteCancel = () => {
    setDeletingUserId(null);
    setDeleteDialogOpen(false);
  };

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
                    {" | "}
                    <button onClick={() => handleDeleteClick(user.Id!)}>
                      削除
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
          <div>
            <Link to={"/user/register"}>新規ユーザー追加</Link>
          </div>

          {/* 削除ダイアログ */}
          {isDeleteDialogOpen && (
            <div className="delete-dialog">
              <p>本当に削除しますか？</p>
              <button onClick={handleDeleteConfirm}>はい</button>
              <button onClick={handleDeleteCancel}>いいえ</button>
            </div>
          )}
        </div>
      ) : (
        <p>ユーザーデータを読み込んでいます...</p>
      )}
    </div>
  );
};

export default UserComponent;
