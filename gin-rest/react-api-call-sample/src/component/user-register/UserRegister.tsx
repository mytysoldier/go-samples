import { useState } from "react";
import { RegisterUser } from "../../api/UserApi";
import { Link } from "react-router-dom";

const UserRegister: React.FC = () => {
  const [name, setName] = useState<string>("");
  const [age, setAge] = useState<number>(0);
  const [isSubmitSuccess, setIsSubmitSuccess] = useState<boolean>(false);

  const handleRegister = async () => {
    const registeredUserData = await RegisterUser({
      Name: name,
      Age: age,
    });

    if (registeredUserData) {
      setIsSubmitSuccess(true);
    }
  };

  return (
    <div>
      <div>
        <h2>User Register</h2>

        <div style={{ marginBottom: "10px" }}>
          <label>Name</label>
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        <div style={{ marginBottom: "10px" }}>
          <label>Age</label>
          <input
            type="number"
            value={age}
            onChange={(e) => setAge(parseInt(e.target.value, 10))}
          />
        </div>

        <div>
          <button onClick={handleRegister} style={{ marginRight: "5px" }}>
            登録
          </button>
          <Link to={"/"}>戻る</Link>
        </div>

        {isSubmitSuccess && <p>登録しました。</p>}
      </div>
    </div>
  );
};

export default UserRegister;
