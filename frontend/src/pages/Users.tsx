import { fetchUsers } from "@/api";
import { useAuth } from "@/auth/AuthContext";
import { User } from "@/types";
import { useEffect, useState } from "react";

const Users = () => {
  const { token } = useAuth();
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    const getUsers = async () => {
      const response = await fetchUsers(token);
      console.log(response);
      setUsers(response.data);
    };
    getUsers();
  }, []);

  return (
    <div className="container mx-auto mt-10">
      <h1 className="text-2xl font-bold">Manage Users</h1>
      <ul className="mt-4 divide-y divide-gray-200">
        {users.map((user) => (
          <li key={user.id} className="flex items-center justify-between py-4">
            <div className="rounded-lg bg-white p-6 shadow-md transition-shadow hover:shadow-lg">
              <h2 className="text-lg font-medium">{user.user_name}</h2>
              <p className="text-sm text-gray-500">
                {user.first_name} {user.last_name}
              </p>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Users;
