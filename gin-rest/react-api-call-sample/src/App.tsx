import React from "react";
import logo from "./logo.svg";
import "./App.css";
import UserComponent from "./component/UserComponent";
import { Route, Router, Routes } from "react-router-dom";
import UserUpdate from "./component/user-update/UserUpdate";

const App: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<UserComponent />} />
      <Route path="user/update/:id" element={<UserUpdate />} />
    </Routes>
  );
};

export default App;
