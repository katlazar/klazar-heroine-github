using HeroesApi.Models;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace HeroesApi.Controllers
{
    [ApiController]
    [Route("api/HeroItems")]
    public class HeroItemsController : ControllerBase
    {
        private readonly HeroContext _context;

        public HeroItemsController(HeroContext context)
        {
            _context = context;
        }

        // GET: api/HeroItems
        [HttpGet]
        public async Task<ActionResult<IEnumerable<HeroItem>>> GetHeroItems()
        {
            return await _context.HeroItems.ToListAsync();
        }

        // GET: api/HeroItems/5
        [HttpGet("{id}")]
        public async Task<ActionResult<HeroItem>> GetHeroItem(long id)
        {
            var heroItem = await _context.HeroItems.FindAsync(id);

            if (heroItem == null)
            {
                return NotFound();
            }

            return heroItem;
        }

        // PUT: api/HeroItems/5
        // To protect from overposting attacks, enable the specific properties you want to bind to, for
        // more details, see https://go.microsoft.com/fwlink/?linkid=2123754.
        [HttpPut("{id}")]
        public async Task<IActionResult> PutHeroItem(long id, HeroItem heroItem)
        {
            if (id != heroItem.Id)
            {
                return BadRequest();
            }

            _context.Entry(heroItem).State = EntityState.Modified;

            try
            {
                await _context.SaveChangesAsync();
            }
            catch (DbUpdateConcurrencyException)
            {
                if (!HeroItemExists(id))
                {
                    return NotFound();
                }
                else
                {
                    throw;
                }
            }

            return NoContent();
        }

        // POST: api/HeroItems
        // To protect from overposting attacks, enable the specific properties you want to bind to, for
        // more details, see https://go.microsoft.com/fwlink/?linkid=2123754.
        [HttpPost]
        public async Task<ActionResult<HeroItem>> PostHeroItem(HeroItem heroItem)
        {
            _context.HeroItems.Add(heroItem);
            await _context.SaveChangesAsync();

            return CreatedAtAction("GetHeroItem", new { id = heroItem.Id }, heroItem);
        }

        // DELETE: api/HeroItems/5
        [HttpDelete("{id}")]
        public async Task<ActionResult<HeroItem>> DeleteHeroItem(long id)
        {
            var heroItem = await _context.HeroItems.FindAsync(id);
            if (heroItem == null)
            {
                return NotFound();
            }

            _context.HeroItems.Remove(heroItem);
            await _context.SaveChangesAsync();

            return heroItem;
        }

        private bool HeroItemExists(long id)
        {
            return _context.HeroItems.Any(e => e.Id == id);
        }
    }
}
