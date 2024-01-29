import { useEffect, useState } from "react";
import User from "../../api/model/User";
import { UpdateUser, fetchUserData } from "../../api/UserApi";
import { Link, useParams } from "react-router-dom";

const UserUpdate: React.FC = () => {
  const { id } = useParams();
  const numericId: number = id ? parseInt(id, 10) : 0;

  const [user, setUser] = useState<User | null>(null);
  const [updatedName, setUpdatedName] = useState<string>("");
  const [updatedAge, setUpdatedAge] = useState<number | undefined>(undefined);
  const [isUpdateSuccess, setIsUpdateSuccess] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchUserData(numericId);
      setUser(data);
    };

    fetchData();
  }, [id]);

  const handleUpdate = async () => {
    if (user) {
      const updatedUserData = await UpdateUser({
        Id: numericId,
        Name: updatedName || user.Name,
        Age: updatedAge !== undefined ? updatedAge : user.Age,
      });

      if (updatedUserData) {
        setIsUpdateSuccess(true);
      }
    }
  };

  return (
    <div>
      {user ? (
        <div>
          <h2>User Update</h2>
          <p>ID: {user.Id}</p>
          <p>Name: {user.Name}</p>
          <p>Age: {user.Age}</p>

          <div style={{ marginBottom: "10px" }}>
            <label>Updated Name:</label>
            <input
              type="text"
              value={updatedName}
              onChange={(e) => setUpdatedName(e.target.value)}
            />
          </div>

          <div style={{ marginBottom: "10px" }}>
            <label>Updated Age:</label>
            <input
              type="number"
              value={updatedAge !== undefined ? updatedAge : ""}
              onChange={(e) => setUpdatedAge(parseInt(e.target.value, 10))}
            />
          </div>

          <div>
            <button onClick={handleUpdate} style={{ marginRight: "5px" }}>
              Update
            </button>
            <Link to={"/"}>戻る</Link>
          </div>

          {isUpdateSuccess && <p>更新しました。</p>}
        </div>
      ) : (
        <p>Loading user data...</p>
      )}
    </div>
  );
};

export default UserUpdate;
