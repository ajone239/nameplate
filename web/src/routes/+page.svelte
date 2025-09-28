<script lang="ts">
    import { onMount } from 'svelte';

    let message = $state('');

    onMount(async () => {
        const res = await fetch('/api/data');

        const { message: message1 } = await res.json();

        message = message1;
    });
</script>

<h1>Welcome to SvelteKit</h1>

<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

<p>The message is: [{message}]</p>
