<script>
    import Status from '$lib/components/Status.svelte';
    import { onMount } from 'svelte';

    const statuses = ['Away', 'Free', 'Busy', 'Meeting', 'Headdown'];

    let interval = 0;

    let selectedOption = $state();
    let status = $state('Away');

    const updateStatus = async () => {
        const res = await fetch('/api/status');
        const { status: new_status } = await res.json();

        status = new_status;
    };

    const updateAll = async () => {
        await updateStatus();
    };

    const submitStatus = async () => {
        const request = {
            status: selectedOption
        };

        const res = await fetch('/api/status', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(request)
        });

        if (res.status != 201) {
            alert(res.body);
            return;
        }

        const { status: new_status } = await res.json();

        status = new_status;
    };

    onMount(async () => {
        await updateAll();
        selectedOption = status;
        interval = setInterval(updateAll, 500);
    });
</script>

<div id="status-setter">
    <h1>Set your status</h1>
    {#each statuses as status}
        <label>
            <input type="radio" name="scoops" value={status} bind:group={selectedOption} />
            {status}
        </label>
    {/each}

    <button disabled={status == selectedOption} onclick={submitStatus}>Submit</button>

    <Status {status} />
</div>

<style>
    #status-setter {
        display: grid;
        flex-direction: column;
    }
</style>
