<div class="min-h-screen min-w-screen flex justify-center items-center">
    <div class="h-fit w-fit border-2 rounded-lg p-6 flex flex-col gap-4">
        <h1>Create a new project</h1>
        <form class="flex flex-col gap-2" action="/api/project" method="post">
            <div class="flex flex-col gap-1">
                <label for="name">Name</label>
                <input type="text" id="name" name="name" required>
            </div>
            <div class="flex flex-col gap-1">
                <label for="description">Description</label>
                <textarea id="description" name="description" required></textarea>
            </div>
            <button class="create-button" type="submit">Create</button>
        </form>
    </div>
</div>
