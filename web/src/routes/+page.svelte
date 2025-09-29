<script lang="ts">
    import { onMount } from 'svelte';

    let message = $state('');

    onMount(async () => {
        const res = await fetch('/api/data');

        const { message: message1 } = await res.json();

        message = message1;
    });
</script>

<div class="m-5">
    <div>
        <h1 class="text-5xl italic underline">Austin Jones</h1>

        <h4 class="text-m text-right italic">Software Engineer</h4>
    </div>
    <p>
        Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation
    </p>

    <p>The message is: [{message}]</p>
</div>
