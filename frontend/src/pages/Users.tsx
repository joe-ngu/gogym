import { fetchUsers } from "@/api";
import { useAuth } from "@/auth/AuthContext";
import UserForm from "@/components/UserForm";
import { User } from "@/types";
import { useEffect, useState } from "react";

const Users = () => {
  const { token } = useAuth();
  const [users, setUsers] = useState<User[]>([]);
  const [editingUser, setEditingUser] = useState<User | null>(null);

  useEffect(() => {
    const getUsers = async () => {
      const response = await fetchUsers(token);
      setUsers(response.data);
    };
    getUsers();
  }, []);

  const handleCreateUser = () => {
    console.log("Create user");
  };

  const handleSave = () => {
    setEditingUser(null);
    const getUsers = async () => {
      const response = await fetchUsers(token);
      setUsers(response.data);
    };
    getUsers();
  };

  return (
    <div className="container mx-auto mt-10">
      <h1 className="text-2xl font-bold">Manage Users</h1>
      <div className="mt-8">
        {editingUser ? (
          <UserForm
            token={token}
            userId={editingUser.id}
            initialData={editingUser}
            onSave={handleSave}
          />
        ) : (
          <button
            onClick={handleCreateUser}
            className="mb-4 inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700"
          >
            Create New User
          </button>
        )}
      </div>
      <ul className="divide-y divide-gray-200 mt-4">
        {users.map((user) => (
          <li key={user.id} className="py-4 flex justify-between items-center">
            <div>
              <h2 className="text-lg font-medium">{user.name}</h2>
              <p className="text-sm text-gray-500">{user.email}</p>
            </div>
            <button
              onClick={() => setEditingUser(user)}
              className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm hover:bg-indigo-700"
            >
              Edit
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Users;
