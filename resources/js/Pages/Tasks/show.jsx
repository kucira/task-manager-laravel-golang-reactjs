import React from 'react';
import { useForm, router, Form , Head} from '@inertiajs/react';

export default function Show({ task }) {
  const { data, put, processing, errors } = useForm({
    name: task.name,
  });

  const handleSubmit = (e) => {
    e.preventDefault();

    put(`/tasks/${task.id}`, {
      onSuccess: () => router.get('/tasks'),
    });
  };

  return (
    <div className="p-6 max-w-xl mx-auto">
    <Head title="Task Manager - Edit" />
      <h1 className="text-2xl font-bold mb-4">Edit Task</h1>

      <Form method="PUT"
        onSuccess={() => router.get('/tasks')}
        >
        <label className="block mb-2 font-medium">Task Name</label>

        <input
          type="text"
        //   value={data.name}
        defaultValue={task.name}
          name="name"
        //   onChange={(e) => setData('name', e.target.value)}
          className="w-full border p-2 rounded"
        />

        {errors.name && (
          <p className="text-red-500 text-sm mt-1">{errors.name}</p>
        )}

        <div className="mt-4 flex gap-2">
          <button
            type="submit"
            disabled={processing}
            className="bg-green-600 text-white px-4 py-2 rounded"
          >
            Save
          </button>

          <button
            type="button"
            className="bg-gray-400 px-4 py-2 rounded"
            onClick={() => router.get('/tasks')}
          >
            Cancel
          </button>
        </div>
      </Form>
    </div>
  );
}
