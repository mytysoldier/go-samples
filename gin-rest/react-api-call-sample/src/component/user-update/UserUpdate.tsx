import { useEffect, useState } from "react";
import "./UserComponent.css";
import User from "../../api/model/User";
import { fetchUserData } from "../../api/UserApi";

const UserUpdate: React.FC<{ id: number }> = ({ id }) => {
  const [user, setUser] = useState<User | null>(null);
  const [updatedName, setUpdatedName] = useState<string>("");
  const [updatedAge, setUpdatedAge] = useState<number | undefined>(undefined);
  const [isUpdateSuccess, setIsUpdateSuccess] = useState<boolean>(false);

  useEffect(() => {
    const fetchData = async () => {
      const data = await fetchUserData(id);
      setUser(data);
    };

    fetchData();
  }, [id]);

  const handleUpdate = async () => {
    if (user) {
      const updatedUserData = await UpdateUser(id, {
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

          <label>
            Updated Name:
            <input
              type="text"
              value={updatedName}
              onChange={(e) => setUpdatedName(e.target.value)}
            />
          </label>

          <label>
            Updated Age:
            <input
              type="number"
              value={updatedAge !== undefined ? updatedAge : ""}
              onChange={(e) => setUpdatedAge(parseInt(e.target.value, 10))}
            />
          </label>

          <button onClick={handleUpdate}>Update</button>

          {isUpdateSuccess && <p>更新しました。</p>}
        </div>
      ) : (
        <p>Loading user data...</p>
      )}
    </div>
  );
};

export default UserUpdate;
