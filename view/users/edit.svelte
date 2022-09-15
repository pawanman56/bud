<script>
  export let user = {}
</script>

<h1>Edit User</h1>

<form method="post" action={`/users/${user.id || 0}`}>
  <input type="hidden" name="_method" value="patch" />
  <!-- Add input fields here -->
  <input type="submit" value="Update User" />
</form>

<br />

<a href={`/users`}>Back</a>
<span> | </span>
<a href={`/users/${user.id || 0}`}>Show User</a>
