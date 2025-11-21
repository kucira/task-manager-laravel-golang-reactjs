<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Task;
use Inertia\Inertia;

class TaskController extends Controller
{
    // homepage index of task page
    public function index() {
        $tasks = Task::orderBy('id', 'desc')->get();
        return Inertia::render(component: 'Tasks/Index', props: ['tasks'=> $tasks]);
    }

    // detail with id
    public function show(Task $task) {
        return Inertia::render(component: 'Tasks/show', props: ['task' => $task]);
    }

    // post data
    public function store(Request $request) {
        \Log::info(message: 'TaskController@store called', context: ['request' => $request->all()]);
        $request->validate(rules: ['name' => 'required']);
        Task::create(attributes: ['name' => $request->name]);
        return redirect()->back();
    }

    // put data
    public function update(Request $request, Task $task) {
        $request->validate(rules: ['name' => 'required']);
        $task->update(attributes: ['name' => $request->name]);
        return redirect()->back();
    }

    // delete
    public function destroy(Task $task) {
        $task->delete();
        return redirect()->back();
    }
}
