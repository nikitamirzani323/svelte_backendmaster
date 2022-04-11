<script>
    import Input_custom from '../../components/Input.svelte' 

    export let pasaran_id = "";
    export let pasaran_nmpasarantogel_field = "";
    export let pasaran_urlpasaran_field = "";
    export let pasaran_pasarandiundi_field = "";
    export let pasaran_jamtutup_field = "";
    export let pasaran_jamjadwal_field = "";
    export let pasaran_jamopen_field = "";
    export let pasaran_status_field = "";

    let pasaran_urlpasaran_error = "";
    let pasaran_pasarandiundi_error = "";
    let pasaran_jamtutup_error = "";
    let pasaran_jamjadwal_error = "";
    let pasaran_jamopen_error = "";
    let pasaran_status_error = "";
</script>

<div class="flex w-full gap-1 mt-2">
    <div class="grid grid-cols-2 gap-2 w-4/5 ">
        <Input_custom
            input_autofocus={false}
            input_required={true}
            input_enabled={true}
            input_tipe="text"
            input_maxlength_text="30"
            bind:value={pasaran_urlpasaran_field}
            input_id="pasaran_urlpasaran_field"
            input_placeholder="Situs"/>
        {#if pasaran_urlpasaran_error != ""}
            <small class="text-pink-600 text-[11px]">{pasaran_urlpasaran_error}</small>
        {/if}
        <Input_custom
            input_autofocus={false}
            input_required={true}
            input_enabled={true}
            input_tipe="time"
            bind:value={pasaran_jamtutup_field}
            input_id="pasaran_jamtutup_field"
            input_placeholder="Tutup"/>
        {#if pasaran_jamtutup_error != ""}
            <small class="text-pink-600 text-[11px]">{pasaran_jamtutup_error}</small>
        {/if}
        <Input_custom
            input_autofocus={false}
            input_required={true}
            input_enabled={true}
            input_tipe="text"
            input_maxlength_text="30"
            bind:value={pasaran_pasarandiundi_field}
            input_id="pasaran_pasarandiundi_field"
            input_placeholder="Diundi"/>
        {#if pasaran_pasarandiundi_error != ""}
            <small class="text-pink-600 text-[11px]">{pasaran_pasarandiundi_error}</small>
        {/if}
        <Input_custom
            input_autofocus={false}
            input_required={true}
            input_enabled={true}
            input_tipe="time"
            bind:value={pasaran_jamjadwal_field}
            input_id="pasaran_jamjadwal_field"
            input_placeholder="Jadwal"/>
        {#if pasaran_jamjadwal_error != ""}
            <small class="text-pink-600 text-[11px]">{pasaran_jamjadwal_error}</small>
        {/if}
        <div></div>
        <Input_custom
            input_autofocus={false}
            input_required={true}
            input_enabled={true}
            input_tipe="time"
            bind:value={pasaran_jamopen_field}
            input_id="pasaran_jamopen_field"
            input_placeholder="Open"/>
        {#if pasaran_jamopen_error != ""}
            <small class="text-pink-600 text-[11px]">{pasaran_jamopen_error}</small>
        {/if}
        <div></div>
        <select
            bind:value={pasaran_status_field}
            class="select select-bordered w-full focus:ring-0 focus:outline-none rounded-md">
            <option disabled selected value="">--Pilih Status--</option>
            <option value="Y">ACTIVE</option>
            <option value="N">DEACTIVE</option>
        </select>
        <div class="col-span-2">
            <button class="btn btn-primary btn-block">Submit</button>
        </div>
    </div>
    <div class="w-full "></div>
</div>