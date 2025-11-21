import React from "react";
import { Head, useForm, Form, Link } from "@inertiajs/react";

export default function Task({ tasks }) {
    const { delete: destroy } = useForm();

    const handleDelete = (id) => {
        return function (event) {
            event.preventDefault();
            destroy(`/tasks/${id}`);
        };
    };

    return (
        <div className="p-6 max-w-xl mx-auto">
            <Head title="Task Manager" />
            <h1 className="text-2xl font-bold mb-4">Task Manager</h1>

            <Form action="/tasks" method="POST" className="flex gap-2 mb-4">
                <input
                    type="text"
                    className="border rounded px-3 py-2 flex-1"
                    placeholder="Enter new task"
                    name="name"
                />
                <button
                    type="submit"
                    className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
                >
                    Add
                </button>
            </Form>

            <ul>
                {tasks.map((task) => (
                    <li
                        key={task.id}
                        className="flex justify-between items-center py-2 border-b"
                    >
                        <span>{task.name}</span>
                        <div className="flex justify-end items-center gap-2">
                            <Link href={`/tasks/${task.id}`} className="text-green-600 hover:text-green-800">
                                Edit
                            </Link>
                            <button
                                onClick={handleDelete(task.id)}
                                className="text-red-600 hover:text-red-800"
                            >
                                Delete
                            </button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
}
